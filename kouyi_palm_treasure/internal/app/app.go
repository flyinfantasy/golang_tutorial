// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/minyou08042/kouyi_palm_treasure/pkg/config"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/login"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go.uber.org/dig"
)

type dserver struct {
	router *gin.Engine
	cont   *dig.Container
	logger logger.LogInfoFormat
}

func NewServer(e *gin.Engine, c *dig.Container, l logger.LogInfoFormat) *dserver {
	return &dserver{
		router: e,
		cont:   c,
		logger: l,
	}
}

func (ds *dserver) SetupDB() error {
	var db *gorm.DB

	if err := ds.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}

	db.Exec("SET search_path TO public")
	/*db.AutoMigrate(
		&pigeon.Pigeon{},
	)*/
	//db.Table("pigeons").Create(&pigeon.Pigeon{})
	//db.AutoMigrate(&user.User{})
	//db.Table("public.pigeon").AutoMigrate(&pigeon.Pigeon{})
	db.AutoMigrate(&login.Login{})
	db.AutoMigrate(&pigeon.Pigeon{})
	db.AutoMigrate(&ptstatus.Ptstatus{})
	return nil
}

func (ds *dserver) Start() error {
	var cfg *config.Config
	if err := ds.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return ds.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
