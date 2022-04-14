package reqTypes

import (
	"github.com/pkg/errors"
)

type Delegate struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Amount           string `json:"amount"`
}

func (d *Delegate) Validate() error {
	switch {
	case d.DelegatorAddress == "":
		return errors.New("empty delegator address")
	case d.ValidatorAddress == "":
		return errors.New("empty validator address")
	case d.Amount == "":
		return errors.New("empty amount")
	default:
		return nil
	}
}

type ReDelegate struct {
	DelegatorAddress    string `json:"delegator_address"`
	ValidatorSrcAddress string `json:"validator_src_address"`
	ValidatorDstAddress string `json:"validator_dst_address"`
	Amount              string `json:"amount"`
}

func (r *ReDelegate) Validate() error {
	switch {
	case r.DelegatorAddress == "":
		return errors.New("empty delegator address")
	case r.ValidatorSrcAddress == "":
		return errors.New("empty validator source address")
	case r.ValidatorDstAddress == "":
		return errors.New("empty validator destination address")
	case r.Amount == "":
		return errors.New("empty amount")
	default:
		return nil
	}
}

type UnDelegate struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Amount           string `json:"amount"`
}

func (u *UnDelegate) Validate() error {
	switch {
	case u.DelegatorAddress == "":
		return errors.New("empty delegator address")
	case u.ValidatorAddress == "":
		return errors.New("empty validator address")
	case u.Amount == "":
		return errors.New("empty amount")
	default:
		return nil
	}
}

type CreateValidator struct {
	Description       description     `json:"description"`
	Commission        commissionRates `json:"commission"`
	MinSelfDelegation string          `json:"min_self_delegation"`
	DelegatorAddress  string          `json:"delegator_address,omitempty" `
	ValidatorAddress  string          `json:"validator_address,omitempty"`
	Pubkey            string          `json:"pubkey,omitempty"`
	Value             string          `json:"value"`
}

func (c *CreateValidator) Validate() error {
	switch {
	case c.Pubkey == "":
		return errors.New("pub_key is null")
	case c.MinSelfDelegation == "":
		return errors.New("min_self_delegation is empty")
	case c.Value == "":
		return errors.New("value is empty")
	case c.DelegatorAddress == "":
		return errors.New("delegator_address is empty")
	default:
		return nil
	}
}

type EditValidator struct {
	Description       description `json:"description"`
	ValidatorAddress  string      `json:"validator_address,omitempty" yaml:"address"`
	CommissionRate    string      `json:"commission_rate,omitempty"`
	MinSelfDelegation string      `json:"min_self_delegation,omitempty"`
}

func (e *EditValidator) Validate() error {
	switch {
	case e.ValidatorAddress == "":
		return errors.New("validator_address is empty")
	default:
		return nil
	}
}

type description struct {
	Moniker         string `json:"moniker,omitempty"`
	Identity        string `json:"identity,omitempty"`
	Website         string `json:"website,omitempty"`
	SecurityContact string `json:"security_contact,omitempty"`
	Details         string `json:"details,omitempty"`
}

type commissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}
