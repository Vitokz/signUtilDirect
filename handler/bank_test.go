package handler

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/Vitokz/signUtilDirect/config"
	"github.com/Vitokz/signUtilDirect/internal/wallet"
	"github.com/cosmos/btcutil/bech32"
	sdkHd "github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkBech32 "github.com/cosmos/cosmos-sdk/types/bech32"
	sdkAuthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_Send(t *testing.T) {
	cfg := config.Parse()
	hnd, err := New(cfg)
	assert.NoError(t, err)

	txBase := "CpIBCo0BChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEm0KK2V0aG0xM2s0MHQ3ZGhhdzd3dW44dWdoNWpsbTJ5bnBwa3E5ZmpudnR6enISK2V0aG0xeDQzbGh0cHBmMGs4MjhjbTA2M3V5MnVsOXhqdG03OWpqcDVqd3gaEQoHYXBob3RvbhIGMTAwMDAwEgASaQpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAtDvXG8rFMYmakCTqgemc4Oj0/Yq1NMjcS5IZe3JF1OBEgQKAggBGC4SFQoPCgdhcGhvdG9uEgQ1MDAwEMCaDBpAcATxdHYWq+8A5vTWnuZ0ONlQSYz/9Emm3X9WSKGaCuI5sEL1VNRUm87CsxWdXFZaX4DBdBLZnAL5gb45dc0crQ=="
	txBytes, err := base64.StdEncoding.DecodeString(txBase)
	assert.NoError(t, err)

	tx, err := hnd.EncConf.TxConfig.TxDecoder()(txBytes)
	assert.NoError(t, err)

	sigTx, _ := tx.(sdkAuthsigning.SigVerifiableTx)

	pubkeys, err := sigTx.GetPubKeys()
	assert.NoError(t, err)

	t.Log(sdk.AccAddress(pubkeys[0].Address()).String())

	//
}

func TestRec(t *testing.T) {
	hexAddress := "BDF8AC2D157358EB7AB17CF97701D83F7F1CD516"

	acc, err := sdk.AccAddressFromHex(hexAddress)
	assert.NoError(t, err)

	hrp, data, err := bech32.Decode(acc.String(), 1023)
	assert.NoError(t, err)

	dec := sdkBech32.DecodeAndConvert
	_ = dec
	t.Log(hrp, data)

	t.Log(acc.String())

	//mnemonic := "spot flush switch era payment family aerobic talk balcony ugly orient marine"
	//params, err := sdkHd.NewParamsFromPath(wallet.HdPath)
	//assert.NoError(t, err)
	//
	//derivedPriv, err := wallet.AlgoEth.Derive()(mnemonic, "", params.String())
	//assert.NoError(t, err)
	//
	//privKey := wallet.AlgoEth.Generate()(derivedPriv)
	//pubkey := privKey.PubKey()
	//
	//t.Log(privKey.Bytes())
	//t.Log(pubkey.Bytes())
	//t.Log(sdk.AccAddress(pubkey.Address()).String())
	//
	//wal := wallet.Wallet{
	//	PrivKey: privKey,
	//	PubKey:  privKey.PubKey(),
	//	Address: pubkey.Address(),
	//}
	//
	//encConfig := ethEncoding.MakeConfig(ethApp.ModuleBasics)
	//
	//hdnl := handler{
	//	Wallet:  &wal,
	//	EncConf: &encConfig,
	//}
	//
	//bz, err := hdnl.Send(context.Background(), &reqTypes.Request{
	//	Msg: &reqTypes.Send{
	//		Amount:      "100000aphoton",
	//		FromAddress: "0xbdf8ac2d157358eb7ab17cf97701d83f7f1cd516",
	//		ToAddress:   "0x8DaAF5F9B7ebBcee4CFc45E92FeD449843601532",
	//	},
	//	Params: reqTypes.Params{
	//		AccountParams: reqTypes.AccountParams{AccountNumber: 845, Sequence: 78},
	//		ChainID:       "worknet_20220112-1",
	//		Fees:          "5000aphoton",
	//		GasWanted:     200000,
	//	},
	//})
	//
	//t.Log(base64.StdEncoding.EncodeToString(bz))
	//
	//t.Log(hex.EncodeToString(pubkey.Address()))
	//t.Log(hex.EncodeToString(privKey.Bytes()))
	//address := sdk.AccAddress(privKey.PubKey().Address())

	//t.Log(bz)
}

func TestBIP(t *testing.T) {
	mnemonic := "shy guitar unknown negative pretty visit please mixed vague wrist require cash finish warrior mimic dragon empty coconut recall wheat flight reveal genuine strong"

	params, err := sdkHd.NewParamsFromPath(wallet.HdPath)
	assert.NoError(t, err)

	derivedPriv, err := wallet.AlgoEth.Derive()(mnemonic, "", params.String())
	assert.NoError(t, err)

	privKey := wallet.AlgoEth.Generate()(derivedPriv)
	t.Log(hex.EncodeToString(privKey.Bytes()))
	address := sdk.AccAddress(privKey.PubKey().Address())

	adr, err := sdk.AccAddressFromBech32("ethm1hpm2nmgw9vsdayn6evsx5vcpju7upfjk3vhjhv")
	assert.NoError(t, err)
	assert.Equal(t, address.String(), adr.String())

	t.Log(address)
	t.Log(adr)
}
