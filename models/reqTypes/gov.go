package reqTypes

import (
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/pkg/errors"
)

type SubmitProposal struct {
	InitialDeposit string `json:"initial_deposit"`
	Proposer       string `json:"proposer,omitempty"`
}

func (s *SubmitProposal) Validate() error {
	switch {
	case s.Proposer == "":
		return errors.New("proposer is empty")
	default:
		return nil
	}
}

type CancelSoftwareUpgradeProposal struct {
	SubmitProposal
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *CancelSoftwareUpgradeProposal) Validate() error {
	err := c.SubmitProposal.Validate()
	switch {
	case err != nil:
		return err
	case c.Title == "":
		return errors.New("title is empty")
	case c.Description == "":
		return errors.New("description is empty")
	default:
		return nil
	}
}

type CommunityPoolSpendProposal struct {
	SubmitProposal
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Recipient   string `json:"recipient,omitempty"`
	Amount      string `json:"amount"`
}

func (c *CommunityPoolSpendProposal) Validate() error {
	err := c.SubmitProposal.Validate()
	switch {
	case err != nil:
		return err
	case c.Title == "":
		return errors.New("title is empty")
	case c.Description == "":
		return errors.New("description is empty")
	case c.Recipient == "":
		return errors.New("recipient is empty")
	case c.Amount == "":
		return errors.New("amount is empty")
	default:
		return nil
	}
}

type ParameterChangeProposal struct {
	SubmitProposal
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Changes     []ParamChange `json:"changes"`
}

func (pcp *ParameterChangeProposal) Validate() error {
	err := pcp.SubmitProposal.Validate()
	switch {
	case err != nil:
		return err
	case pcp.Title == "":
		return errors.New("title is empty")
	case pcp.Description == "":
		return errors.New("description is empty")
	case len(pcp.Changes) > 0:
		for _, v := range pcp.Changes {
			if err = v.validate(); err != nil {
				return err
			}
		}
	default:
		return nil
	}

	return nil
}

func (pcp *ParameterChangeProposal) ToParamChanges() []proposal.ParamChange {
	res := make([]proposal.ParamChange, len(pcp.Changes))
	for i, pc := range pcp.Changes {
		res[i] = pc.ToParamChange()
	}
	return res
}

type SoftwareUpgradeProposal struct {
	SubmitProposal
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Plan        Plan   `json:"plan"`
}

func (s *SoftwareUpgradeProposal) Validate() error {
	err := s.SubmitProposal.Validate()
	switch {
	case err != nil:
		return err
	case s.Title == "":
		return errors.New("title is empty")
	case s.Description == "":
		return errors.New("description is empty")
	default:
		return s.Plan.validate()
	}
}

type Deposit struct {
	ProposalID uint64 `json:"proposal_id"`
	Depositor  string `json:"depositor,omitempty"`
	Amount     string `json:"amount"`
}

func (d *Deposit) Validate() error {
	switch {
	case d.ProposalID == 0:
		return errors.New("proposalID is empty")
	case d.Depositor == "":
		return errors.New("depositor is empty")
	case d.Amount == "":
		return errors.New("amount is empty")
	default:
		return nil
	}
}

type Vote struct {
	ProposalID uint64 `json:"proposal_id"`
	Voter      string `json:"voter,omitempty"`
	Option     string `json:"option,omitempty"`
}

func (v *Vote) Validate() error {
	switch {
	case v.ProposalID == 0:
		return errors.New("proposalID is empty")
	case v.Voter == "":
		return errors.New("voter is empty")
	case v.Option == "":
		return errors.New("option is empty")
	default:
		return nil
	}
}

type VoteWeighted struct {
	ProposalID uint64   `json:"proposal_id,omitempty"`
	Voter      string   `json:"voter,omitempty"`
	Options    []string `json:"options"`
}

func (vm *VoteWeighted) Validate() error {
	switch {
	case vm.ProposalID == 0:
		return errors.New("proposalID is empty")
	case vm.Voter == "":
		return errors.New("voter is empty")
	case len(vm.Options) == 0:
		return errors.New("options is empty")
	default:
		return nil
	}
}

//// ------------------------
type Plan struct {
	Name   string `json:"name,omitempty"`
	Height int64  `json:"height,omitempty"`
	Info   string `json:"info,omitempty"`
}

func (p *Plan) validate() error {
	switch {
	case p.Name == "":
		return errors.New("plan name is empty")
	case p.Height == 0:
		return errors.New("plan height is empty")
	case p.Info == "":
		return errors.New("plan info is empty")
	default:
		return nil
	}
}

type ParamChange struct {
	Subspace string `json:"subspace,omitempty"`
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
}

func (pc *ParamChange) validate() error {
	switch {
	case pc.Key == "":
		return errors.New("key is empty in paramChange")
	case pc.Subspace == "":
		return errors.New("subspace is empty in paramChange")
	case pc.Value == "":
		return errors.New("value is empty in paramChange")
	default:
		return nil
	}
}

func (pc *ParamChange) ToParamChange() proposal.ParamChange {
	return proposal.NewParamChange(pc.Subspace, pc.Key, pc.Value)
}
