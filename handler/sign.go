package handler

import (
	"context"
	"github.com/Vitokz/signUtilDirect/internal/txFactory"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkSigning "github.com/cosmos/cosmos-sdk/types/tx/signing"
	sdkAuthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
)

func (h *handler) Sign(ctx context.Context, req *reqTypes.UnsignedTxRequest) ([]byte, error) {
	var (
		txBytes = req.GetTx()
		params  = req.GetParams()
	)

	tx, err := h.EncConf.TxConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode unsigned tx")
	}

	if err = checkFromAndWalletAddrEqual(h.Wallet.Address, params.From); err != nil {
		return nil, err
	}

	factory, err := h.createTxFactory(ctx, params)
	if err != nil {
		return nil, err
	}

	txBuilder, err := h.EncConf.TxConfig.WrapTxBuilder(tx)
	if err != nil {
		return nil, err
	}

	err = h.signTx(factory, txBuilder, false)
	if err != nil {
		return nil, err
	}

	if params.PrintSignatureOnly {
		return marshalSignatureJSON(h.EncConf.TxConfig, txBuilder, params.PrintSignatureOnly)
	}

	return factory.TxConfig.TxEncoder()(txBuilder.GetTx())
}

func (h *handler) signTx(txf *txFactory.Factory, txBuilder sdkClient.TxBuilder, overwriteSig bool) (err error) {
	signMode := txf.SignMode
	if signMode == sdkSigning.SignMode_SIGN_MODE_UNSPECIFIED {
		// use the SignModeHandler's default mode if unspecified
		signMode = txf.TxConfig.SignModeHandler().DefaultMode()
	}

	if err = checkMultipleSigners(signMode, txBuilder.GetTx()); err != nil {
		return err
	}

	pubKey := h.Wallet.PubKey
	signerData := sdkAuthsigning.SignerData{
		ChainID:       txf.ChainID,
		AccountNumber: txf.AccountNumber,
		Sequence:      txf.Sequence,
	}
	sigData := sdkSigning.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := sdkSigning.SignatureV2{
		PubKey:   pubKey,
		Data:     &sigData,
		Sequence: txf.GetSequence(),
	}

	var prevSignatures []sdkSigning.SignatureV2
	if !overwriteSig {
		prevSignatures, err = txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			return err
		}
	}

	if err = txBuilder.SetSignatures(sig); err != nil {
		return err
	}

	bytesToSign, err := txf.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return err
	}

	sigBytes, _, err := h.Wallet.Sign(bytesToSign)

	sigData = sdkSigning.SingleSignatureData{
		SignMode:  signMode,
		Signature: sigBytes,
	}
	sig = sdkSigning.SignatureV2{
		PubKey:   pubKey,
		Data:     &sigData,
		Sequence: txf.GetSequence(),
	}

	if overwriteSig {
		return txBuilder.SetSignatures(sig)
	}
	prevSignatures = append(prevSignatures, sig)
	return txBuilder.SetSignatures(prevSignatures...)
}

func checkMultipleSigners(mode sdkSigning.SignMode, tx sdkAuthsigning.Tx) error {
	if mode == sdkSigning.SignMode_SIGN_MODE_DIRECT &&
		len(tx.GetSigners()) > 1 {
		return errors.Wrap(sdkerrors.ErrNotSupported, "Signing in DIRECT mode is only supported for transactions with one signer only")
	}
	return nil
}

func marshalSignatureJSON(txConfig sdkClient.TxConfig, txBldr sdkClient.TxBuilder, signatureOnly bool) ([]byte, error) {
	parsedTx := txBldr.GetTx()
	if signatureOnly {
		sigs, err := parsedTx.GetSignaturesV2()
		if err != nil {
			return nil, err
		}
		return txConfig.MarshalSignatureJSON(sigs)
	}

	return txConfig.TxJSONEncoder()(parsedTx)
}

func checkFromAndWalletAddrEqual(wallet crypto.Address, from string) error {
	if from == "" {
		return errors.New("param 'from' is empty")
	}

	fromAdr, err := hexAddressToAccAddress(from)
	if err != nil {
		return err
	}

	if sdk.AccAddress(wallet).String() != fromAdr.String() {
		return errors.New("from address and signer address are not equal")
	}

	return nil
}
