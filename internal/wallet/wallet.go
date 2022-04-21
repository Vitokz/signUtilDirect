package wallet

import (
	"github.com/Vitokz/signUtilDirect/config"
	sdkHd "github.com/cosmos/cosmos-sdk/crypto/hd"
	cCryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
	tCryptotypes "github.com/tendermint/tendermint/crypto"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	eCryptoHd "github.com/tharsis/ethermint/crypto/hd"
)

var (
	AlgoEth = eCryptoHd.EthSecp256k1
	HdPath  = "m/44'/60'/0'/0/0"
)

type Wallet struct {
	Address  tCryptotypes.Address
	PubKey   cCryptoTypes.PubKey
	PrivKey  cCryptoTypes.LedgerPrivKey
	Sequence int
}

func NewWalletFromMnemonic(mnemonic string) (*Wallet, error) {
	params, err := sdkHd.NewParamsFromPath(HdPath)
	if err != nil {
		return nil, err
	}

	derivedPriv, err := AlgoEth.Derive()(mnemonic, "", params.String())
	if err != nil {
		return nil, err
	}

	privKey := AlgoEth.Generate()(derivedPriv)
	pubKey := privKey.PubKey()
	address := pubKey.Address()

	return &Wallet{
		PrivKey: privKey,
		PubKey:  pubKey,
		Address: address,
	}, nil
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
