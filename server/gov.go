package server

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) vote(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.Vote{},
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

	tx, err := s.Handler.Vote(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) deposit(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.Deposit{},
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

	tx, err := s.Handler.Deposit(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) voteWeighted(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.VoteWeighted{},
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

	tx, err := s.Handler.VoteWeighted(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) cancelSoftwareUpgradeProposal(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.CancelSoftwareUpgradeProposal{},
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

	tx, err := s.Handler.CancelSoftwareUpgradeProposal(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) communityPoolSpendProposal(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.CommunityPoolSpendProposal{},
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

	tx, err := s.Handler.CommunityPoolSpendProposal(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) parameterChangeProposal(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.ParameterChangeProposal{},
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

	tx, err := s.Handler.ParameterChangeProposal(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) softwareUpgradeProposal(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.SoftwareUpgradeProposal{},
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

	tx, err := s.Handler.SoftwareUpgradeProposal(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}
