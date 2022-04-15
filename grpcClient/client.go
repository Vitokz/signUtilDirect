package grpcClient

import (
	"github.com/Vitokz/signUtilDirect/config"
	sdkAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	sdkDistributionTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"google.golang.org/grpc"
)

type Client struct {
	AuthClient         sdkAuthTypes.QueryClient
	DistributionClient sdkDistributionTypes.QueryClient

	cons []*grpc.ClientConn
}

func NewClient(cfg config.Config) (*Client, error) {
	var (
		addr = cfg.GetGrpcAddress()
		err  error
	)

	grpc.WithInsecure()

	aClient, aConn, err := newAuthClient(addr)
	if err != nil {
		return nil, err
	}

	dstClient, dstConn, err := newDistributionClient(addr)
	if err != nil {
		return nil, err
	}

	client := Client{
		AuthClient:         aClient,
		DistributionClient: dstClient,
	}

	client.cons = append(client.cons,
		aConn,
		dstConn,
	)

	return &client, nil
}

func newAuthClient(addr string) (sdkAuthTypes.QueryClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := sdkAuthTypes.NewQueryClient(conn)

	return client, nil, nil
}

func newDistributionClient(addr string) (sdkDistributionTypes.QueryClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := sdkDistributionTypes.NewQueryClient(conn)

	return client, nil, nil
}

func (c *Client) StopAllCons() {
	for _, v := range c.cons {
		v.Close()
	}
}
