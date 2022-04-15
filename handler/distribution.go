package handler

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkDistributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/pkg/errors"
)

func (h *handler) FundCommunityPool(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertFundCommunityPoolMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) SetWithdrawAddress(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertSetWithdrawAddressMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) WithdrawDelegatorReward(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msgs, err := convertWithdrawDelegatorRewardMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msgs...)
}

func (h *handler) WithdrawAllDelegatorRewards(ctx context.Context, req *reqTypes.Request) (map[int][]byte, error) {
	var (
		params = req.GetParams()
	)

	msgs, err := h.buildWithdrawAllDelegatorRewardsMsgs(ctx, req)
	if err != nil {
		return nil, err
	}

	txs := make(map[int][]byte)
	for i, msg := range msgs {
		txs[i], err = h.buildAndSignTx(ctx, params, msg)
		if err != nil {
			return nil, err
		}
	}

	return txs, nil
}

func (h *handler) buildWithdrawAllDelegatorRewardsMsgs(ctx context.Context, req *reqTypes.Request) ([]sdk.Msg, error) {
	warReq, ok := req.GetMsg().(*reqTypes.WithdrawAllDelegatorRewards)
	if !ok {
		return nil, errors.New("failed to convert withdrawAllDelegatorRewards structure from interface")
	}

	delAddr, err := hexAddressToAccAddress(warReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	valRes, err := h.Client.DistributionClient.DelegatorValidators(ctx, &sdkDistributionTypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: delAddr.String(),
	})

	validators := valRes.Validators

	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		val, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			return nil, err
		}

		msg := sdkDistributionTypes.NewMsgWithdrawDelegatorReward(delAddr, val)
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func convertFundCommunityPoolMsg(req *reqTypes.Request) (*sdkDistributionTypes.MsgFundCommunityPool, error) {
	sendReq, ok := req.GetMsg().(*reqTypes.FundCommunityPool)
	if !ok {
		return nil, errors.New("failed to convert fundCommunityPool structure from interface")
	}

	coins, err := sdk.ParseCoinsNormalized(sendReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	depositorAddr, err := hexAddressToAccAddress(sendReq.Depositor)
	if err != nil {
		return nil, err
	}

	return sdkDistributionTypes.NewMsgFundCommunityPool(coins, depositorAddr), nil
}

func convertSetWithdrawAddressMsg(req *reqTypes.Request) (*sdkDistributionTypes.MsgSetWithdrawAddress, error) {
	swaReq, ok := req.GetMsg().(*reqTypes.SetWithdrawAddress)
	if !ok {
		return nil, errors.New("failed to convert setWithdrawAddress structure from interface")
	}

	withdrawAddr, err := hexAddressToAccAddress(swaReq.WithdrawAddress)
	if err != nil {
		return nil, err
	}

	delAddr, err := hexAddressToAccAddress(swaReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	return sdkDistributionTypes.NewMsgSetWithdrawAddress(delAddr, withdrawAddr), nil
}

func convertWithdrawDelegatorRewardMsg(req *reqTypes.Request) ([]sdk.Msg, error) {
	wdrReq, ok := req.GetMsg().(*reqTypes.WithdrawDelegatorReward)
	if !ok {
		return nil, errors.New("failed to convert withdrawDelegatorReward structure from interface")
	}

	valAddr, err := hexAddressToValAddress(wdrReq.ValidatorAddress)
	if err != nil {
		return nil, err
	}

	delAddr, err := hexAddressToAccAddress(wdrReq.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	msgs := []sdk.Msg{sdkDistributionTypes.NewMsgWithdrawDelegatorReward(delAddr, valAddr)}

	if commission := req.GetParams().Commission; commission {
		msgs = append(msgs, sdkDistributionTypes.NewMsgWithdrawValidatorCommission(valAddr))
	}

	return msgs, nil
}
