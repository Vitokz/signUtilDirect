package handler

import (
	"context"
	"fmt"
	"github.com/Vitokz/signUtilDirect/internal/txFactory"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
)

func (h *handler) buildAndSignTx(ctx context.Context, params reqTypes.Params, msg ...sdk.Msg) ([]byte, error) {
	factory, err := h.createTxFactory(ctx, params)
	if err != nil {
		return nil, err
	}

	tx, err := factory.BuildUnsignedTx(msg...)
	if err != nil {
		return nil, err
	}

	if params.FeeAccount != "" {
		granter, err := hexAddressToAccAddress(params.FeeAccount)
		if err != nil {
			fmt.Println(err)
		}
		tx.SetFeeGranter(granter)
	}

	err = h.signTx(factory, tx, false)
	if err != nil {
		return nil, err
	}

	txBytes, err := factory.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, err
}

func hexAddressToAccAddress(hex string) (sdk.AccAddress, error) {
	address, err := sdk.AccAddressFromHex(hex[2:])
	if err != nil {
		return nil, errors.Wrap(err, "failed convert hex acc address")
	}

	return address, nil
}

func hexAddressToValAddress(hex string) (sdk.ValAddress, error) {
	address, err := sdk.ValAddressFromHex(hex[2:])
	if err != nil {
		return nil, errors.Wrap(err, "failed convert hex val address")
	}

	return address, nil
}

func (h *handler) createTxFactory(ctx context.Context, params reqTypes.Params) (*txFactory.Factory, error) {
	factory, err := txFactory.NewTxFactory(params, h.EncConf.TxConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create tx factory")
	}

	err = h.takeAccountParams(ctx, factory)
	if err != nil {
		return nil, err
	}

	return factory, nil
}

func (h *handler) takeAccountParams(ctx context.Context, builder *txFactory.Factory) error {
	accAny, err := h.Client.AuthClient.Account(ctx, &sdkAuthTypes.QueryAccountRequest{
		Address: sdk.AccAddress(h.Wallet.Address.Bytes()).String(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to take account data")
	}

	var acc sdkAuthTypes.AccountI
	err = h.EncConf.Marshaler.UnpackAny(accAny.Account, &acc)
	if err != nil {
		return errors.Wrap(err, "failed to unpack account any data")
	}

	if builder.AccountNumber == 0 || builder.AccountNumber < acc.GetAccountNumber() {
		builder.AccountNumber = acc.GetAccountNumber()
	}
	if builder.Sequence == 0 || builder.Sequence < acc.GetSequence() {
		builder.Sequence = acc.GetSequence()
	}

	return nil
}
