package main

import (
	"github.com/Vitokz/signUtilDirect/config"
	"github.com/Vitokz/signUtilDirect/handler"
	"github.com/Vitokz/signUtilDirect/server"
	"github.com/pkg/errors"
	"log"
)

func main() {
	cfg := config.Parse()

	hdlr, err := handler.New(cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create handler: "))
	}
	defer hdlr.Client.StopAllCons()

	serv := server.New(hdlr)

	serv.Start(cfg)
}
