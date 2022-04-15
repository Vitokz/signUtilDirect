package reqTypes

import "github.com/pkg/errors"

type Send struct {
	FromAddress string `json:"from_address,omitempty"`
	ToAddress   string `json:"to_address,omitempty"`
	Amount      string `json:"amount"`
}

func (s *Send) Validate() error {
	switch {
	case s.FromAddress == "":
		return errors.New("from_address is empty")
	case s.ToAddress == "":
		return errors.New("to_address is empty")
	case s.Amount == "":
		return errors.New("amount is empty")
	default:
		return nil
	}
}
