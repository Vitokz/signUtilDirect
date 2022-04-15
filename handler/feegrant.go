package handler

import (
	"context"
	"fmt"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkFeegrant "github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/pkg/errors"
	"time"
)

func (h *handler) Grant(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertGrantMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func (h *handler) Revoke(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertRevokeMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertGrantMsg(req *reqTypes.Request) (*sdkFeegrant.MsgGrantAllowance, error) {
	grantReq, ok := req.GetMsg().(*reqTypes.Grant)
	if !ok {
		return nil, errors.New("failed to convert grant structure from interface")
	}

	granterAddr, err := hexAddressToAccAddress(grantReq.Granter)
	if err != nil {
		return nil, err
	}

	granteeAddr, err := hexAddressToAccAddress(grantReq.Grantee)
	if err != nil {
		return nil, err
	}

	limit, err := sdk.ParseCoinsNormalized(grantReq.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse limit")
	}

	basic := sdkFeegrant.BasicAllowance{
		SpendLimit: limit,
	}

	exp := grantReq.Expiration
	var expiresAtTime time.Time
	if exp != "" {
		expiresAtTime, err = time.Parse(time.RFC3339, exp)
		if err != nil {
			return nil, err
		}
		basic.Expiration = &expiresAtTime
	}

	var grant sdkFeegrant.FeeAllowanceI
	grant = &basic

	period := int64(grantReq.Period)
	periodLimitStr := grantReq.PeriodLimit

	if period > 0 || periodLimitStr != "" {
		periodLimit, err := sdk.ParseCoinsNormalized(periodLimitStr)
		if err != nil {
			return nil, err
		}

		if period <= 0 {
			return nil, fmt.Errorf("period clock was not set")
		}

		if periodLimit == nil {
			return nil, fmt.Errorf("period limit was not set")
		}

		periodReset := getPeriodReset(period)
		if exp != "" && periodReset.Sub(expiresAtTime) > 0 {
			return nil, fmt.Errorf("period (%d) cannot reset after expiration (%v)", period, exp)
		}

		periodic := sdkFeegrant.PeriodicAllowance{
			Basic:            basic,
			Period:           getPeriod(period),
			PeriodReset:      getPeriodReset(period),
			PeriodSpendLimit: periodLimit,
			PeriodCanSpend:   periodLimit,
		}

		grant = &periodic
	}

	allowedMsgs := grantReq.AllowedMsgs
	if len(allowedMsgs) > 0 {
		grant, err = sdkFeegrant.NewAllowedMsgAllowance(grant, allowedMsgs)
		if err != nil {
			return nil, err
		}
	}

	return sdkFeegrant.NewMsgGrantAllowance(grant, granterAddr, granteeAddr)
}

func convertRevokeMsg(req *reqTypes.Request) (*sdkFeegrant.MsgRevokeAllowance, error) {
	revokeReq, ok := req.GetMsg().(*reqTypes.Revoke)
	if !ok {
		return nil, errors.New("failed to convert revoke structure from interface")
	}

	granterAddr, err := hexAddressToAccAddress(revokeReq.Granter)
	if err != nil {
		return nil, err
	}

	granteeAddr, err := hexAddressToAccAddress(revokeReq.Grantee)
	if err != nil {
		return nil, err
	}

	msg := sdkFeegrant.NewMsgRevokeAllowance(granterAddr, granteeAddr)
	return &msg, nil
}

func getPeriodReset(duration int64) time.Time {
	return time.Now().Add(getPeriod(duration))
}

func getPeriod(duration int64) time.Duration {
	return time.Duration(duration) * time.Second
}
