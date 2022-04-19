package handler

import (
	"encoding/base64"
	"github.com/Vitokz/signUtilDirect/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkAuthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Send(t *testing.T) {
	cfg := config.Parse()
	hd, err := New(cfg)
	assert.NoError(t, err)

	txBase := "CpIBCo0BChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEm0KK2V0aG0xM2s0MHQ3ZGhhdzd3dW44dWdoNWpsbTJ5bnBwa3E5ZmpudnR6enISK2V0aG0xeDQzbGh0cHBmMGs4MjhjbTA2M3V5MnVsOXhqdG03OWpqcDVqd3gaEQoHYXBob3RvbhIGMTAwMDAwEgASaQpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAtDvXG8rFMYmakCTqgemc4Oj0/Yq1NMjcS5IZe3JF1OBEgQKAggBGC4SFQoPCgdhcGhvdG9uEgQ1MDAwEMCaDBpAcATxdHYWq+8A5vTWnuZ0ONlQSYz/9Emm3X9WSKGaCuI5sEL1VNRUm87CsxWdXFZaX4DBdBLZnAL5gb45dc0crQ=="
	txBytes, err := base64.StdEncoding.DecodeString(txBase)
	assert.NoError(t, err)

	tx, err := hd.EncConf.TxConfig.TxDecoder()(txBytes)
	assert.NoError(t, err)

	sigTx, _ := tx.(sdkAuthsigning.SigVerifiableTx)

	pubkeys, err := sigTx.GetPubKeys()
	assert.NoError(t, err)

	t.Log(sdk.AccAddress(pubkeys[0].Address()).String())

	//
	//msg2 := sdkTx.Fee{}
	////msg := sdkBankTypes.MsgSend{}
	//t.Log(proto.MessageName(&msg2))
}
