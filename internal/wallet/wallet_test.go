package wallet

import (
	"github.com/Vitokz/signUtilDirect/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	"os"
	"testing"
)

func TestNewWallet(t *testing.T) {
	assert.NoError(t, os.Setenv("SIGN_APP_VAL_PRIV_KEY", "bcf2441a0087d841cd9e4a44fd37c7a09b87cf29ef7e1042a11797a049884d28"))
	cfg := config.Parse()
	wal := NewWallet(cfg)

	require.Equal(t, wal.Address.String(), "57617E6A47B6D6ADC7B773C08B4626ACA0904AE2")
	require.Equal(t, wal.PubKey.String(), "EthPubKeySecp256k1{03D6133ABFBCFDEFE0B479BF7303549D3589CD1926D3186FA6A640DD2CAA29F4BE}")
}

func TestWallet_Sign(t *testing.T) {
	privKey, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)

	wal := Wallet{PrivKey: privKey}

	msg := []byte("hello world")
	sigHash := crypto.Keccak256Hash(msg)
	expectedSig, err := secp256k1.Sign(sigHash.Bytes(), privKey.Bytes())
	require.NoError(t, err)

	sig, _, err := wal.Sign(sigHash.Bytes())
	require.NoError(t, err)
	require.Equal(t, expectedSig, sig)
}
