package reqTypes

type Msg interface {
	Validate() error
}

type Request struct {
	Msg    Msg    `json:"msg"`
	Params Params `json:"params"`
}

func (r *Request) GetParams() Params {
	return r.Params
}

func (r *Request) GetMsg() Msg {
	return r.Msg
}

type UnsignedTxRequest struct {
	Tx     []byte `json:"tx"`
	Params Params `json:"params"`
}

func (u *UnsignedTxRequest) GetTx() []byte {
	return u.Tx
}

func (u *UnsignedTxRequest) GetParams() Params {
	return u.Params
}
