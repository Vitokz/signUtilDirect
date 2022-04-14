package server

import (
	"github.com/Vitokz/signUtilDirect/config"
	"github.com/Vitokz/signUtilDirect/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
	//logger ...
	Handler handler.Handler
}

func New(hdlr handler.Handler) *Server {
	newServ := &Server{
		Router:  newRouter(),
		Handler: hdlr,
	}

	newServ.routing()

	return newServ
}

func newRouter() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger())

	return router
}

func (s *Server) Start(cfg config.Config) {
	s.routing()

	s.Router.Logger.Fatal(s.Router.Start(":" + cfg.GetPort()))
}

func (s *Server) routing() {
	gTx := s.Router.Group("/tx")

	gTx.POST("/staking/delegate", s.Delegate)
	gTx.POST("/staking/redelegate", s.ReDelegate)
	gTx.POST("/staking/unbond", s.UnDelegate)
	gTx.POST("/staking/create_validator", s.CreateValidator)
	gTx.POST("/staking/edit_validator", s.EditValidator)

	gTx.POST("/sign", s.Sign)
	_ = gTx
}
