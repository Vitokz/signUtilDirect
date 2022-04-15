package server

import (
	"context"
	"github.com/Vitokz/signUtilDirect/models/reqTypes"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) fundCommunityPool(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.FundCommunityPool{},
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

	tx, err := s.Handler.FundCommunityPool(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) setWithdrawAddress(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.SetWithdrawAddress{},
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

	tx, err := s.Handler.SetWithdrawAddress(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) withdrawDelegatorReward(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.WithdrawDelegatorReward{},
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

	tx, err := s.Handler.WithdrawDelegatorReward(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, reqTypes.Response{Tx: tx})
}

func (s *Server) withdrawAllDelegatorRewards(c echo.Context) error {
	var (
		request = &reqTypes.Request{
			Msg: &reqTypes.WithdrawAllDelegatorRewards{},
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

	txs, err := s.Handler.WithdrawAllDelegatorRewards(context.Background(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	respTxs := make([][]byte, len(txs))
	for _, tx := range txs {
		respTxs = append(respTxs, tx)
	}

	return c.JSON(http.StatusOK, reqTypes.BatchTxResponse{Txs: respTxs})
}
