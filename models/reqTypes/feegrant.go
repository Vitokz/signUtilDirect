package reqTypes

import "errors"

type Grant struct {
	Granter     string   `json:"granter"`
	Grantee     string   `json:"grantee"`
	Limit       string   `json:"limit"`
	Expiration  string   `json:"expiration,omitempty"` //"2006-01-02T15:04:05Z07:00" RFC3339
	Period      int      `json:"period,omitempty"`
	PeriodLimit string   `json:"periodLimit,omitempty"`
	AllowedMsgs []string `json:"allowedMsgs,omitempty1"`
}

func (g *Grant) Validate() error {
	switch {
	case g.Granter == "":
		return errors.New("granter is empty")
	case g.Grantee == "":
		return errors.New("grantee is empty")
	case g.Limit == "":
		return errors.New("limit is empty")
	default:
		return nil
	}
}

type Revoke struct {
	Granter string `json:"granter"`
	Grantee string `json:"grantee"`
}

func (r *Revoke) Validate() error {
	switch {
	case r.Granter == "":
		return errors.New("granter is empty")
	case r.Grantee == "":
		return errors.New("grantee is empty")
	default:
		return nil
	}
}
