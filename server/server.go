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

	// -- staking
	stakg := gTx.Group("/staking")
	stakg.POST("/delegate", s.delegate)
	stakg.POST("/redelegate", s.reDelegate)
	stakg.POST("/unbond", s.unDelegate)
	stakg.POST("/create_validator", s.createValidator)
	stakg.POST("/edit_validator", s.editValidator)

	// -- bank
	bankg := gTx.Group("/bank")
	bankg.POST("/send", s.send)

	// -- distribution
	dst := gTx.Group("/distribution")
	dst.POST("/fund_community_pool", s.fundCommunityPool)
	dst.POST("/set_withdraw_address", s.setWithdrawAddress)
	dst.POST("/withdraw_delegator_reward", s.withdrawDelegatorReward)
	dst.POST("/withdraw_all_delegator_reward", s.withdrawAllDelegatorRewards)

	// -- feegrant
	grant := gTx.Group("/feegrant")
	grant.POST("/grant", s.grant)
	grant.POST("/revoke", s.revoke)

	// -- gov
	gov := gTx.Group("/gov")
	gov.POST("/deposit", s.deposit)
	gov.POST("/vote", s.vote)
	gov.POST("/vote_weighted", s.voteWeighted)

	govsp := gov.Group("submit_proposal")
	govsp.POST("/cancel_software_upgrade", s.cancelSoftwareUpgradeProposal)
	govsp.POST("/community_pool_spend", s.communityPoolSpendProposal)
	govsp.POST("/param_change", s.parameterChangeProposal)
	govsp.POST("software_upgrade", s.softwareUpgradeProposal)

	// --
	gTx.POST("/sign", s.sign)
	_ = gTx
}
