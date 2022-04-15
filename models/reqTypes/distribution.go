package reqTypes

import "github.com/pkg/errors"

type FundCommunityPool struct {
	Amount    string `json:"amount"`
	Depositor string `json:"depositor,omitempty"`
}

func (f *FundCommunityPool) Validate() error {
	switch {
	case f.Amount == "":
		return errors.New("amount is empty")
	case f.Depositor == "":
		return errors.New("depositor is empty")
	default:
		return nil
	}
}

type SetWithdrawAddress struct {
	DelegatorAddress string `json:"delegator_address,omitempty"`
	WithdrawAddress  string `json:"withdraw_address,omitempty"`
}

func (s *SetWithdrawAddress) Validate() error {
	switch {
	case s.WithdrawAddress == "":
		return errors.New("empty withdraw_address")
	case s.DelegatorAddress == "":
		return errors.New("empty delegator_address")
	default:
		return nil
	}
}

type WithdrawDelegatorReward struct {
	DelegatorAddress string `json:"delegator_address,omitempty"`
	ValidatorAddress string `json:"validator_address,omitempty"`
}

func (w *WithdrawDelegatorReward) Validate() error {
	switch {
	case w.ValidatorAddress == "":
		return errors.New("empty validator_address")
	case w.DelegatorAddress == "":
		return errors.New("empty delegator_address")
	default:
		return nil
	}
}

type WithdrawAllDelegatorRewards struct {
	DelegatorAddress string `json:"delegator_address,omitempty"`
}

func (w *WithdrawAllDelegatorRewards) Validate() error {
	switch {
	case w.DelegatorAddress == "":
		return errors.New("empty delegator_address")
	default:
		return nil
	}
}
