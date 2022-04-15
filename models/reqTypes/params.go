package reqTypes

type AccountParams struct {
	AccountNumber uint64 `json:"account_number,omitempty"`
	Sequence      uint64 `json:"sequence,omitempty"`
}

type Params struct {
	// -- Factory Params
	AccountParams
	FeeAccount    string  `json:"fee_account,omitempty"`
	ChainID       string  `json:"chainID,omitempty"`
	From          string  `json:"from,omitempty"`
	GasPrices     string  `json:"gas_prices,omitempty"`
	GasWanted     uint64  `json:"gas_wanted,omitempty"`
	Node          string  `json:"node,omitempty"`
	GasAdjustment float64 `json:"gas_adjustment,omitempty"`
	Fees          string  `json:"fees,omitempty"`
	SignMode      string  `json:"sign_mode,omitempty"`

	// -- Only Sign Params
	PrintSignatureOnly bool `json:"print_signature_only"`

	// -- Distribution Params
	Commission bool `json:"commission"`

	// -- Other Params
	Overwrite bool `json:"overwrite"`
}
