package handler

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkDistribTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	sdkGovUtils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	sdkGovTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	sdkUpgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/pkg/errors"
	"strings"
)

func (h *handler) CancelSoftwareUpgradeProposal(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertCancelSoftwareUpgradeProposal(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertCancelSoftwareUpgradeProposal(req *reqTypes.Request) (*sdkGovTypes.MsgSubmitProposal, error) {
	propReq, ok := req.GetMsg().(*reqTypes.CancelSoftwareUpgradeProposal)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	deposit, err := sdk.ParseCoinsNormalized(propReq.InitialDeposit)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse deposit")
	}

	proposer, err := hexAddressToAccAddress(propReq.Proposer)
	if err != nil {
		return nil, err
	}

	content := sdkUpgradeTypes.NewCancelSoftwareUpgradeProposal(propReq.Title, propReq.Description)

	return sdkGovTypes.NewMsgSubmitProposal(content, deposit, proposer)
}

func (h *handler) CommunityPoolSpendProposal(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertCommunityPoolSpendProposal(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertCommunityPoolSpendProposal(req *reqTypes.Request) (*sdkGovTypes.MsgSubmitProposal, error) {
	propReq, ok := req.GetMsg().(*reqTypes.CommunityPoolSpendProposal)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	deposit, err := sdk.ParseCoinsNormalized(propReq.InitialDeposit)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse deposit")
	}

	amount, err := sdk.ParseCoinsNormalized(propReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	proposer, err := hexAddressToAccAddress(propReq.Proposer)
	if err != nil {
		return nil, err
	}

	recipient, err := hexAddressToAccAddress(propReq.Recipient)
	if err != nil {
		return nil, err
	}

	content := sdkDistribTypes.NewCommunityPoolSpendProposal(propReq.Title, propReq.Description, recipient, amount)

	return sdkGovTypes.NewMsgSubmitProposal(content, deposit, proposer)
}

func (h *handler) ParameterChangeProposal(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertParameterChangeProposal(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertParameterChangeProposal(req *reqTypes.Request) (*sdkGovTypes.MsgSubmitProposal, error) {
	propReq, ok := req.GetMsg().(*reqTypes.ParameterChangeProposal)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	proposer, err := hexAddressToAccAddress(propReq.Proposer)
	if err != nil {
		return nil, err
	}

	deposit, err := sdk.ParseCoinsNormalized(propReq.InitialDeposit)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse deposit")
	}

	content := paramproposal.NewParameterChangeProposal(
		propReq.Title, propReq.Description, propReq.ToParamChanges(),
	)

	return sdkGovTypes.NewMsgSubmitProposal(content, deposit, proposer)
}

func (h *handler) SoftwareUpgradeProposal(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertSoftwareUpgradeProposal(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertSoftwareUpgradeProposal(req *reqTypes.Request) (*sdkGovTypes.MsgSubmitProposal, error) {
	propReq, ok := req.GetMsg().(*reqTypes.SoftwareUpgradeProposal)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	proposer, err := hexAddressToAccAddress(propReq.Proposer)
	if err != nil {
		return nil, err
	}

	deposit, err := sdk.ParseCoinsNormalized(propReq.InitialDeposit)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse deposit")
	}

	content := sdkUpgradeTypes.NewSoftwareUpgradeProposal(
		propReq.Title, propReq.Description, sdkUpgradeTypes.Plan{
			Name:   propReq.Plan.Name,
			Height: propReq.Plan.Height,
			Info:   propReq.Plan.Info,
		})

	return sdkGovTypes.NewMsgSubmitProposal(content, deposit, proposer)
}

func (h *handler) Deposit(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertDepositMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertDepositMsg(req *reqTypes.Request) (*sdkGovTypes.MsgDeposit, error) {
	propReq, ok := req.GetMsg().(*reqTypes.Deposit)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	depositor, err := hexAddressToAccAddress(propReq.Depositor)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinsNormalized(propReq.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed parse amount")
	}

	msg := sdkGovTypes.NewMsgDeposit(depositor, propReq.ProposalID, amount)

	return msg, nil
}

func (h *handler) Vote(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertVoteMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertVoteMsg(req *reqTypes.Request) (*sdkGovTypes.MsgVote, error) {
	propReq, ok := req.GetMsg().(*reqTypes.Vote)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	voter, err := hexAddressToAccAddress(propReq.Voter)
	if err != nil {
		return nil, err
	}

	byteVoteOption, err := sdkGovTypes.VoteOptionFromString(sdkGovUtils.NormalizeVoteOption(propReq.Option))
	if err != nil {
		return nil, err
	}

	msg := sdkGovTypes.NewMsgVote(voter, propReq.ProposalID, byteVoteOption)

	return msg, nil
}

func (h *handler) VoteWeighted(ctx context.Context, req *reqTypes.Request) ([]byte, error) {
	var (
		params = req.GetParams()
	)

	msg, err := convertVoteWeightedMsg(req)
	if err != nil {
		return nil, err
	}

	return h.buildAndSignTx(ctx, params, msg)
}

func convertVoteWeightedMsg(req *reqTypes.Request) (*sdkGovTypes.MsgVoteWeighted, error) {
	propReq, ok := req.GetMsg().(*reqTypes.VoteWeighted)
	if !ok {
		return nil, errors.New("failed to convert cancelSoftwareUpgradeProposal structure from interface")
	}

	voter, err := hexAddressToAccAddress(propReq.Voter)
	if err != nil {
		return nil, err
	}

	options, err := sdkGovTypes.WeightedVoteOptionsFromString(sdkGovUtils.NormalizeWeightedVoteOptions(strings.Join(propReq.Options, ",")))
	if err != nil {
		return nil, err
	}

	msg := sdkGovTypes.NewMsgVoteWeighted(voter, propReq.ProposalID, options)

	return msg, nil
}
