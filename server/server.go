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

	sg := gTx.Group("/staking")
	sg.POST("/delegate", s.delegate)
	sg.POST("/redelegate", s.reDelegate)
	sg.POST("/unbond", s.unDelegate)
	sg.POST("/create_validator", s.createValidator)
	sg.POST("/edit_validator", s.editValidator)

	bg := gTx.Group("/bank")
	bg.POST("/send", s.send)

	gTx.POST("/sign", s.sign)
	_ = gTx
}
