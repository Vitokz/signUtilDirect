package server

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) delegate(c echo.Context) error {
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

func (s *Server) reDelegate(c echo.Context) error {
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

func (s *Server) unDelegate(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.UnDelegate{},
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

	tx, err := s.Handler.UnDelegate(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) createValidator(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.CreateValidator{},
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

	tx, err := s.Handler.CreateValidator(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) editValidator(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.CreateValidator{},
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

	tx, err := s.Handler.EditValidator(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}
