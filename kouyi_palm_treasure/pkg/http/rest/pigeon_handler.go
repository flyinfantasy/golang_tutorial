package rest

import (
	"net/http"

	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type sentpigeonCtrl struct {
	log logger.LogInfoFormat
	svc pigeon.Service
}

/*type ClubandDate struct {
	Pigeon_club									string `json:"pigeon_club"`
	Induction_date						  string `json:"induction_date"`
}*/

func NewPigeonCtrl(log logger.LogInfoFormat, svc pigeon.Service) *sentpigeonCtrl {
	return &sentpigeonCtrl{log, svc}
}

func (p *sentpigeonCtrl) GetAll(ctx *gin.Context) {
	pigeons, err := p.svc.GetAll()
	if len(pigeons) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, pigeons)
}

func (p *sentpigeonCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
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

func (p *sentpigeonCtrl) GetByDate(ctx *gin.Context) {
	induction_date := ctx.Param("induction_date")
	pigeons, err := p.svc.GetByDate(induction_date)
	if pigeons == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, pigeons)
}

func (p *sentpigeonCtrl) Store(ctx *gin.Context) {
	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pigeon.Status_code == "02" {
		println(pigeon.Status_code)
		p.svc.Store(&pigeon)
		ctx.Status(http.StatusCreated)
	} else if pigeon.Status_code == "01" {
		return
	}
}

func (p *sentpigeonCtrl) Update(ctx *gin.Context) {
	Uid := ctx.Param("Uid")
	if _, err := uuid.FromString(Uid); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pigeon.Uid = Uid
	p.svc.Update(&pigeon)
	ctx.Status(http.StatusOK)
}

func (p *sentpigeonCtrl) WebStore(ctx *gin.Context) {
	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.svc.WebStore(&pigeon)
	ctx.Status(http.StatusCreated)
}

func (p *sentpigeonCtrl) WebUpdate(ctx *gin.Context) {
	Uid := ctx.Param("Uid")
	if _, err := uuid.FromString(Uid); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pigeon.Uid = Uid
	p.svc.WebUpdate(&pigeon)
	ctx.Status(http.StatusOK)
}

func (p *sentpigeonCtrl) PtStore(ctx *gin.Context) {
	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.svc.PtStore(&pigeon)
	ctx.Status(http.StatusCreated)
}

func (p *sentpigeonCtrl) PtUpdate(ctx *gin.Context) {
	Uid := ctx.Param("Uid")
	if _, err := uuid.FromString(Uid); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var pigeon pigeon.Pigeon
	if err := ctx.ShouldBindJSON(&pigeon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error by": err.Error()})
		return
	}
	pigeon.Uid = Uid
	p.svc.PtUpdate(&pigeon)
	ctx.Status(http.StatusOK)
}

func (p *sentpigeonCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("Uid")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	p.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}

/*
	func (p *sentpigeonCtrl) GetByDate(ctx *gin.Context) {
		induction_date := ctx.Param("induction_date")
		pigeons, err := p.svc.GetByDate(induction_date)
		if pigeons == nil || err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, pigeons)
	}
*/
func (p *sentpigeonCtrl) GetByClub(ctx *gin.Context) {
	pigeon_club := ctx.Param("pigeon_club")
	pigeons, err := p.svc.GetByClub(pigeon_club)
	if pigeons == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, pigeons)
}

func (p *sentpigeonCtrl) GetByClubandDate(ctx *gin.Context) {
	pigeon_club := ctx.Param("pigeon_club")
	date := ctx.Param("date")
	pigeons, err := p.svc.GetByClubandDate(pigeon_club, date)
	if pigeons == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, pigeons)
}
