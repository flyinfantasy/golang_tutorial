package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/login"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type loginCtrl struct {
	log logger.LogInfoFormat
	svc login.Service
}

func NewLoginCtrl(log logger.LogInfoFormat, svc login.Service) *loginCtrl {
	return &loginCtrl{log, svc}
}

func (l *loginCtrl) GetAll(ctx *gin.Context) {
	logins, err := l.svc.GetAll()
	if len(logins) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, logins)
}

/*
	func (p *sentptstatusCtrl) Store(ctx *gin.Context) {
		var ptstatus ptstatus.Ptstatus
		if err := ctx.ShouldBindJSON(&ptstatus); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p.svc.Store(&ptstatus)
		ctx.Status(http.StatusCreated)
	}
*/
func (l *loginCtrl) Store(ctx *gin.Context) {
	var login login.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	l.svc.Store(&login)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
	})
}

func (l *loginCtrl) Update(ctx *gin.Context) {
	Pigeon_account := ctx.Param("Pigeon_account")
	if _, err := uuid.FromString(Pigeon_account); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var login login.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	login.Pigeon_tel = Pigeon_account
	l.svc.Update(&login)
	ctx.Status(http.StatusOK)
}

func (l *loginCtrl) Delete(ctx *gin.Context) {
	Pigeon_account := ctx.Param("Pigeon_account")
	if _, err := uuid.FromString(Pigeon_account); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	l.svc.Delete(Pigeon_account)
	ctx.Status(http.StatusNoContent)
}

func (l *loginCtrl) Signin(ctx *gin.Context) {
	var login login.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Pigeon_account := ctx.Param("Pigeon_account")
	/*if _, err := uuid.FromString(Pigeon_account); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}*/
	//Pigeon_password := ctx.Param("Pigeon_password")
	/*if _, err := uuid.FromString(Pigeon_password); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}*/
	logina := l.svc.Signin(login.Pigeon_tel, login.Pigeon_loft)
	ctx.JSON(http.StatusOK, gin.H{
		"pigeon_club": logina,
	})
}
