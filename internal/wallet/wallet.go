package wallet

import (
	"github.com/Vitokz/signUtilDirect/config"
	cCryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
	tCryptotypes "github.com/tendermint/tendermint/crypto"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
)

type Wallet struct {
	Address tCryptotypes.Address
	PubKey  cCryptoTypes.PubKey
	PrivKey cCryptoTypes.LedgerPrivKey
	//passwordPrivKey string
}

func NewWallet(cfg config.Config) *Wallet {
	privKey := &ethsecp256k1.PrivKey{
		Key: common.FromHex(cfg.GetPrivKey()),
	}
	pubKey := privKey.PubKey()
	address := pubKey.Address()

	return &Wallet{
		PrivKey: privKey,
		PubKey:  pubKey,
		Address: address,
	}
}

func (w *Wallet) Sign(msg []byte) ([]byte, cCryptoTypes.PubKey, error) {
	sig, err := w.PrivKey.Sign(msg)
	if err != nil {
		return nil, nil, err
	}

	return sig, w.PrivKey.PubKey(), nil
}
