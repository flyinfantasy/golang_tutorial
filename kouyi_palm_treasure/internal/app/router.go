package app

import (
	"github.com/minyou08042/kouyi_palm_treasure/pkg/http/rest"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/login"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"

	"github.com/gin-gonic/gin"
)

func (ds *dserver) MapRoutes() {

	// Group : v1
	apiV1 := ds.router.Group("api/v1")

	ds.healthRoutes(apiV1)
	ds.loginRoutes(apiV1)
	ds.pigeonRoutes(apiV1)
	ds.ptstatusRoutes(apiV1)
	ds.mixInputRoutes(apiV1)
}

func (ds *dserver) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := rest.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}

func (ds *dserver) loginRoutes(api *gin.RouterGroup) {
	var loginSvc login.Service
	ds.cont.Invoke(func(l login.Service) {
		loginSvc = l
	})

	loginRoutes := api.Group("/login")
	{
		f := rest.NewLoginCtrl(ds.logger, loginSvc)
		loginRoutes.POST("/signin/", f.Signin)
		loginRoutes.POST("/", f.Store)
		loginRoutes.GET("/GetAll/", f.GetAll)
		loginRoutes.DELETE("/:id", f.Delete)
		loginRoutes.PUT("/:id", f.Update)
	}
}

func (ds *dserver) pigeonRoutes(api *gin.RouterGroup) {
	pigeonRoutes := api.Group("/pigeons")
	{
		var pigeonSvc pigeon.Service
		ds.cont.Invoke(func(u pigeon.Service) {
			pigeonSvc = u
		})

		pigeo := rest.NewPigeonCtrl(ds.logger, pigeonSvc)

		pigeonRoutes.GET("/", pigeo.GetAll)
		pigeonRoutes.POST("/", pigeo.Store)
		pigeonRoutes.POST("/web/", pigeo.WebStore)
		pigeonRoutes.POST("/pt/", pigeo.PtStore)
		pigeonRoutes.GET("/id/:id", pigeo.GetByID)
		pigeonRoutes.GET("/date/:induction_date", pigeo.GetByDate)
		pigeonRoutes.PUT("/:id", pigeo.Update)
		pigeonRoutes.DELETE("/:id", pigeo.Delete)
		pigeonRoutes.GET("/club/:pigeon_club", pigeo.GetByClub)
		pigeonRoutes.GET("/clubanddate/:pigeon_club/:date", pigeo.GetByClubandDate)
	}
}

func (ds *dserver) ptstatusRoutes(api *gin.RouterGroup) {
	ptstatusRoutes := api.Group("/ptstatuss")
	{
		var ptstatusSvc ptstatus.Service
		ds.cont.Invoke(func(u ptstatus.Service) {
			ptstatusSvc = u
		})

		ptstatu := rest.NewPtstatusCtrl(ds.logger, ptstatusSvc)

		ptstatusRoutes.GET("/", ptstatu.GetAll)
		ptstatusRoutes.POST("/", ptstatu.Store)
		ptstatusRoutes.GET("/id/:id", ptstatu.GetByID)
		ptstatusRoutes.GET("/date/:palm_treasure_date", ptstatu.GetByDate)
		ptstatusRoutes.PUT("/:id", ptstatu.Update)
		ptstatusRoutes.DELETE("/:id", ptstatu.Delete)
		ptstatusRoutes.GET("/club/:pigeon_club", ptstatu.GetByClub)
		ptstatusRoutes.GET("/clubanddate/:pigeon_club/:date", ptstatu.GetByClubandDate)

	}
}

func (ds *dserver) mixInputRoutes(api *gin.RouterGroup) {
	mixinputRoutes := api.Group("/input")
	{
		var ptstatusSvc ptstatus.Service
		ds.cont.Invoke(func(u ptstatus.Service) {
			ptstatusSvc = u
		})

		var pigeonSvc pigeon.Service
		ds.cont.Invoke(func(v pigeon.Service) {
			pigeonSvc = v
		})
		mixinput := rest.NewMixInputCtrl(ds.logger, pigeonSvc, ptstatusSvc)
		mixinputRoutes.POST("/", mixinput.Store)
	}
}
