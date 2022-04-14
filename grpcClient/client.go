package grpcClient

import (
	"github.com/Vitokz/signUtilDirect/config"
	sdkAuthTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc"
)

type Client struct {
	AuthClient sdkAuthTypes.QueryClient

	cons []*grpc.ClientConn
}

func NewClient(cfg config.Config) (*Client, error) {
	var (
		addr   = cfg.GetGrpcAddress()
		err    error
		client = &Client{}
	)

	aClient, aConn, err := newAuthClient(addr)
	if err != nil {
		return client, err
	}

	client.AuthClient = aClient
	client.cons = append(client.cons, aConn)

	return client, nil
}

func newAuthClient(addr string) (sdkAuthTypes.QueryClient, *grpc.ClientConn, error) {

	grpc.WithInsecure()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := sdkAuthTypes.NewQueryClient(conn)

	return client, nil, nil
}

func (c *Client) StopAllCons() {
	for _, v := range c.cons {
		v.Close()
	}
}
