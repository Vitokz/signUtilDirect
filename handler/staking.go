package handler

import (
	"context"
	"fmt"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkCryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkStakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
)

func (h *handler) Delegate(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertDelegateMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) ReDelegate(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertReDelegateMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) UnDelegate(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertUnDelegateMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) CreateValidator(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertCreateValidatorMsg(req, h.EncConf.Marshaler)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) EditValidator(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertEditValidatorMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertDelegateMsg(req *reqTypes.Request) (*sdkStakingTypes.MsgDelegate, error) {
	delegateReq, ok := req.GetMsg().(*reqTypes.Delegate)
	if !ok {
		return nil, errors.New("failed to convert delegate structure from interface")
	}

	coins, err := sdk.ParseCoinNormalized(delegateReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	valAdr, err := hexAddressToValAddress(delegateReq.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	delAdr, err := hexAddressToAccAddress(delegateReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	return sdkStakingTypes.NewMsgDelegate(delAdr, valAdr, coins), nil
}

func convertReDelegateMsg(req *reqTypes.Request) (*sdkStakingTypes.MsgBeginRedelegate, error) {
	reDelegateReq, ok := req.GetMsg().(*reqTypes.ReDelegate)
	if !ok {
		return nil, errors.New("failed to convert reDelegate structure from interface")
	}

	coins, err := sdk.ParseCoinNormalized(reDelegateReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	valSrcAdr, err := hexAddressToValAddress(reDelegateReq.ValidatorSrcAddress)
	if err != nil {
		return nil, err
	}

	valDstAdr, err := hexAddressToValAddress(reDelegateReq.ValidatorDstAddress)
	if err != nil {
		return nil, err
	}

	delAdr, err := hexAddressToAccAddress(reDelegateReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	return sdkStakingTypes.NewMsgBeginRedelegate(delAdr, valSrcAdr, valDstAdr, coins), nil
}

func convertUnDelegateMsg(req *reqTypes.Request) (*sdkStakingTypes.MsgUndelegate, error) {
	unDelegateReq, ok := req.GetMsg().(*reqTypes.Delegate)
	if !ok {
		return nil, errors.New("failed to convert unDelegate structure from interface")
	}

	coins, err := sdk.ParseCoinNormalized(unDelegateReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	valAdr, err := hexAddressToValAddress(unDelegateReq.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	delAdr, err := hexAddressToAccAddress(unDelegateReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	return sdkStakingTypes.NewMsgUndelegate(delAdr, valAdr, coins), nil
}

func convertCreateValidatorMsg(req *reqTypes.Request, marshaller codec.Codec) (*sdkStakingTypes.MsgCreateValidator, error) {
	createValidatorReq, ok := req.GetMsg().(*reqTypes.CreateValidator)
	if !ok {
		return nil, errors.New("failed to convert create structure from interface")
	}

	coins, err := sdk.ParseCoinNormalized(createValidatorReq.Value)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	var pk sdkCryptoTypes.PubKey
	if err = marshaller.UnmarshalInterfaceJSON([]byte(createValidatorReq.Pubkey), &pk); err != nil {
		return nil, err
	}

	valAdr, err := hexAddressToValAddress(createValidatorReq.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	desc := sdkStakingTypes.NewDescription(
		createValidatorReq.Description.Moniker,
		createValidatorReq.Description.Identity,
		createValidatorReq.Description.Website,
		createValidatorReq.Description.SecurityContact,
		createValidatorReq.Description.Details,
	)

	commissionRates, err := buildCommissionRates(
		createValidatorReq.Commission.Rate,
		createValidatorReq.Commission.MaxRate,
		createValidatorReq.Commission.MaxChangeRate)
	if err != nil {
		return nil, err
	}

	minSelfDelegation, ok := sdk.NewIntFromString(createValidatorReq.MinSelfDelegation)
	if !ok {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "minimum self delegation must be a positive integer")
	}

	msg, err := sdkStakingTypes.NewMsgCreateValidator(
		valAdr, pk, coins, desc, commissionRates, minSelfDelegation,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert createValidator msg")
	}

	err = msg.ValidateBasic()
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate new createValidator msg")
	}

	return msg, err
}

func convertEditValidatorMsg(req *reqTypes.Request) (*sdkStakingTypes.MsgEditValidator, error) {
	editValidatorReq, ok := req.GetMsg().(*reqTypes.EditValidator)
	if !ok {
		return nil, errors.New("failed to convert edit validator structure")
	}

	valAdr, err := hexAddressToValAddress(editValidatorReq.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	desc := sdkStakingTypes.NewDescription(
		editValidatorReq.Description.Moniker,
		editValidatorReq.Description.Identity,
		editValidatorReq.Description.Website,
		editValidatorReq.Description.SecurityContact,
		editValidatorReq.Description.Details,
	)

	var newRate *sdk.Dec

	commissionRate := editValidatorReq.CommissionRate
	if commissionRate != "" {
		rate, err := sdk.NewDecFromStr(commissionRate)
		if err != nil {
			return nil, fmt.Errorf("invalid new commission rate: %v", err)
		}

		newRate = &rate
	}

	var newMinSelfDelegation *sdk.Int

	minSelfDelegationString := editValidatorReq.MinSelfDelegation
	if minSelfDelegationString != "" {
		msb, ok := sdk.NewIntFromString(minSelfDelegationString)
		if !ok {
			return nil, errors.New("minimum self delegation must be a positive integer")
		}

		newMinSelfDelegation = &msb
	}

	return sdkStakingTypes.NewMsgEditValidator(valAdr, desc, newRate, newMinSelfDelegation), nil
}

func buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr string) (commission sdkStakingTypes.CommissionRates, err error) {
	if rateStr == "" || maxRateStr == "" || maxChangeRateStr == "" {
		return commission, errors.New("must specify all validator commission parameters")
	}

	rate, err := sdk.NewDecFromStr(rateStr)
	if err != nil {
		return commission, err
	}

	maxRate, err := sdk.NewDecFromStr(maxRateStr)
	if err != nil {
		return commission, err
	}

	maxChangeRate, err := sdk.NewDecFromStr(maxChangeRateStr)
	if err != nil {
		return commission, err
	}

	commission = sdkStakingTypes.NewCommissionRates(rate, maxRate, maxChangeRate)

	return commission, nil
}
