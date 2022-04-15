package reqTypes

type Response struct {
	Tx []byte `json:"tx"`
}

type BatchTxResponse struct {
	Txs [][]byte `json:"txs"`
}
