package handler

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkBankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
)

func (h *handler) Send(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertSendMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, msg, params)
}

func convertSendMsg(req *reqTypes.Request) (*sdkBankTypes.MsgSend, error) {
	sendReq, ok := req.GetMsg().(*reqTypes.Send)
	if !ok {
		return nil, errors.New("failed to convert send structure from interface")
	}

	coins, err := sdk.ParseCoinsNormalized(sendReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	fromAdr, err := hexAddressToAccAddress(sendReq.FromAddress)
	if err != nil {
		return nil, err
	}

	toAdr, err := hexAddressToAccAddress(sendReq.ToAddress)
	if err != nil {
		return nil, err
	}

	return sdkBankTypes.NewMsgSend(fromAdr, toAdr, coins), nil
}
