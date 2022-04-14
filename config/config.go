package config

import "os"

type config struct {
	appPort     string
	grpcAddress string
	wallet      wallet
}

type wallet struct {
	privKey string
}

func Parse() Config {
	cfg := new(config)

	if cfg.appPort = os.Getenv("SIGN_APP_PORT"); cfg.appPort == "" {
		cfg.appPort = "3582"
	}

	if cfg.grpcAddress = os.Getenv("SIGN_APP_GRPC_ADDRESS"); cfg.grpcAddress == "" {
		cfg.grpcAddress = "localhost:9090"
	}

	if cfg.wallet.privKey = os.Getenv("SIGN_APP_VAL_PRIV_KEY"); cfg.wallet.privKey == "" {
		cfg.wallet.privKey = "bcf2441a0087d841cd9e4a44fd37c7a09b87cf29ef7e1042a11797a049884d28"
	}

	return cfg
}

type Config interface {
	GetPort() string
	GetGrpcAddress() string
	GetPrivKey() string
}

func (c *config) GetPort() string {
	return c.appPort
}

func (c *config) GetGrpcAddress() string {
	return c.grpcAddress
}

func (c *config) GetPrivKey() string {
	return c.wallet.privKey
}
