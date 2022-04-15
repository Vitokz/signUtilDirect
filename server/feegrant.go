package server

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) grant(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.Grant{},
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

	tx, err := s.Handler.Grant(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) revoke(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.Revoke{},
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

	tx, err := s.Handler.Revoke(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}
