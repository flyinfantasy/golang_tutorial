package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"
	"net/http"
)

type sentmixInputCtrl struct {
	log  logger.LogInfoFormat
	svcn pigeon.Service
	svcs ptstatus.Service
}

func NewMixInputCtrl(log logger.LogInfoFormat, svcn pigeon.Service, svcs ptstatus.Service) *sentmixInputCtrl {
	return &sentmixInputCtrl{log, svcn, svcs}
}

func (p *sentmixInputCtrl) Store(ctx *gin.Context) {
	var ptstatus ptstatus.Ptstatus
	if err := ctx.ShouldBindBodyWith(&ptstatus, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.svcs.Store(&ptstatus)
	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindBodyWith(&pigeon, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pigeon.Status_code == "02" {
		p.svcn.PtStore(&pigeon)
		ctx.Status(http.StatusCreated)
	}
}
