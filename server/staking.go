package server

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) Delegate(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.Delegate{},
		}
		err error
	)
	err = c.Bind(request)
	if err != nil {
		//h.log.WithFields(logrus.Fields{
		//	"event: ": "registration user",
		//	"err: ":   err,
		//	"time: ":  time.Now(),
		//})

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = request.Msg.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tx, err := s.Handler.Delegate(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) ReDelegate(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.ReDelegate{},
		}
		err error
	)
	err = c.Bind(request)
	if err != nil {
		//h.log.WithFields(logrus.Fields{
		//	"event: ": "registration user",
		//	"err: ":   err,
		//	"time: ":  time.Now(),
		//})

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err = request.Msg.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tx, err := s.Handler.ReDelegate(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}
