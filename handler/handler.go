package handler

import (
	"context"
	mycfg "github.com/Vitokz/signUtilDirect/config"
	"github.com/Vitokz/signUtilDirect/grpcClient"
	"github.com/Vitokz/signUtilDirect/internal/wallet"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	sdkParams "github.com/cosmos/cosmos-sdk/simapp/params"
	ethApp "github.com/tharsis/ethermint/app"
	ethEncoding "github.com/tharsis/ethermint/encoding"
)

type Handler interface {
	Sign(ctx context.Context, req *reqTypes.UnsignedTxRequest) ([]byte, error)

	handlerStakingInt
	handlerBankInt
	handlerDistributionInt
	handlerFeegrant
}

type handler struct {
	EncConf *sdkParams.EncodingConfig
	Client  *grpcClient.Client
	Wallet  *wallet.Wallet
}

func New(cfg mycfg.Config) (*handler, error) {
	encCfg := ethEncoding.MakeConfig(ethApp.ModuleBasics)

	client, err := grpcClient.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	wal := wallet.NewWallet(cfg)

	return &handler{
		EncConf: &encCfg,
		Client:  client,
		Wallet:  wal,
	}, nil
}

type handlerStakingInt interface {
	Delegate(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	ReDelegate(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	UnDelegate(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	CreateValidator(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	EditValidator(ctx context.Context, req *reqTypes.Request) ([]byte, error)
}

type handlerBankInt interface {
	Send(ctx context.Context, req *reqTypes.Request) ([]byte, error)
}

type handlerDistributionInt interface {
	FundCommunityPool(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	SetWithdrawAddress(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	WithdrawDelegatorReward(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	WithdrawAllDelegatorRewards(ctx context.Context, req *reqTypes.Request) (map[int][]byte, error)
}

type handlerFeegrant interface {
	Grant(ctx context.Context, req *reqTypes.Request) ([]byte, error)
	Revoke(ctx context.Context, req *reqTypes.Request) ([]byte, error)
}
