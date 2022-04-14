package txFactory

import (
	"fmt"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/pkg/errors"
)

type Factory struct {
	TxConfig      client.TxConfig
	AccountNumber uint64
	Sequence      uint64
	Gas           uint64
	TimeoutHeight uint64
	GasAdjustment float64
	ChainID       string
	Memo          string
	Fees          sdk.Coins
	GasPrices     sdk.DecCoins
	SignMode      signing.SignMode
}

func (f Factory) GetAccountNumber() uint64   { return f.AccountNumber }
func (f Factory) GetSequence() uint64        { return f.Sequence }
func (f Factory) GetGas() uint64             { return f.Gas }
func (f Factory) GetGasAdjustment() float64  { return f.GasAdjustment }
func (f Factory) GetChainID() string         { return f.ChainID }
func (f Factory) GetMemo() string            { return f.Memo }
func (f Factory) GetFees() sdk.Coins         { return f.Fees }
func (f Factory) GetGasPrices() sdk.DecCoins { return f.GasPrices }
func (f Factory) GetTimeoutHeight() uint64   { return f.TimeoutHeight }

func NewTxFactory(params reqTypes.Params, txConfig client.TxConfig) (*Factory, error) {
	var err error

	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	switch params.SignMode {
	case flags.SignModeDirect:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	case flags.SignModeEIP191:
		signMode = signing.SignMode_SIGN_MODE_EIP_191
	default:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	}

	f := Factory{
		TxConfig:      txConfig,
		Gas:           params.GasWanted,
		GasAdjustment: params.GasAdjustment,
		ChainID:       params.ChainID,
		SignMode:      signMode,
	}

	if f, err = f.WithFees(params.Fees); err != nil {
		return nil, err
	}
	if f, err = f.WithGasPrices(params.GasPrices); err != nil {
		return nil, err
	}

	return &f, nil
}

func (f Factory) WithFees(fees string) (Factory, error) {
	parsedFees, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return f, err
	}

	f.Fees = parsedFees
	return f, nil
}

func (f Factory) WithGasPrices(gasPrices string) (Factory, error) {
	parsedGasPrices, err := sdk.ParseDecCoins(gasPrices)
	if err != nil {
		return f, err
	}

	f.GasPrices = parsedGasPrices
	return f, nil
}

func (f Factory) BuildUnsignedTx(msgs ...sdk.Msg) (client.TxBuilder, error) {
	if f.ChainID == "" {
		return nil, fmt.Errorf("chain ID required but not specified")
	}

	fees := f.Fees

	if !f.GasPrices.IsZero() {
		if !fees.IsZero() {
			return nil, errors.New("cannot provide both fees and gas prices")
		}

		glDec := sdk.NewDec(int64(f.Gas))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		fees = make(sdk.Coins, len(f.GasPrices))

		for i, gp := range f.GasPrices {
			fee := gp.Amount.Mul(glDec)
			fees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}

	tx := f.TxConfig.NewTxBuilder()

	if err := tx.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	tx.SetMemo(f.Memo)
	tx.SetFeeAmount(fees)
	tx.SetGasLimit(f.Gas)
	tx.SetTimeoutHeight(f.GetTimeoutHeight())

	return tx, nil
}
