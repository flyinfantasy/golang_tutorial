package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"
	uuid "github.com/satori/go.uuid"
)

type sentptstatusCtrl struct {
	log logger.LogInfoFormat
	svc ptstatus.Service
}

func NewPtstatusCtrl(log logger.LogInfoFormat, svc ptstatus.Service) *sentptstatusCtrl {
	return &sentptstatusCtrl{log, svc}
}

func (p *sentptstatusCtrl) GetAll(ctx *gin.Context) {
	ptstatuss, err := p.svc.GetAll()
	if len(ptstatuss) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, ptstatuss)
}

func (p *sentptstatusCtrl) GetByDate(ctx *gin.Context) {
	palm_treasure_date := ctx.Param("palm_treasure_date")
	ptstatuss, err := p.svc.GetByDate(palm_treasure_date)
	if ptstatuss == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, ptstatuss)
}

func (p *sentptstatusCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("Pigeon_club_pigeon_loft")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := p.svc.GetByID(id)
	if user == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (p *sentptstatusCtrl) Store(ctx *gin.Context) {
	var ptstatus ptstatus.Ptstatus
	if err := ctx.ShouldBindJSON(&ptstatus); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.svc.Store(&ptstatus)
	ctx.Status(http.StatusCreated)
}

func (p *sentptstatusCtrl) Update(ctx *gin.Context) {
	UID_Table := ctx.Param("uID_Table")
	if _, err := uuid.FromString(UID_Table); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var ptstatus ptstatus.Ptstatus
	if err := ctx.ShouldBindJSON(&ptstatus); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ptstatus.UID_Table = UID_Table
	p.svc.Update(&ptstatus)
	ctx.Status(http.StatusOK)
}

func (p *sentptstatusCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("uid")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	p.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}
func (p *sentptstatusCtrl) GetByClub(ctx *gin.Context) {
	pigeon_club := ctx.Param("pigeon_club")
	ptstatus, err := p.svc.GetByClub(pigeon_club)
	if ptstatus == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, ptstatus)
}

func (p *sentptstatusCtrl) GetByClubandDate(ctx *gin.Context) {
	pigeon_club := ctx.Param("pigeon_club")
	date := ctx.Param("date")
	ptstatus, err := p.svc.GetByClubandDate(pigeon_club, date)
	if ptstatus == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, ptstatus)
}
