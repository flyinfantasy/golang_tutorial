package dig

import (
	"github.com/minyou08042/kouyi_palm_treasure/pkg/config"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/login"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/pigeon"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/ptstatus"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/storage"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/storage/orm"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// logger
	container.Provide(logger.NewLogger)

	// login
	container.Provide(orm.NewLoginRepo)
	container.Provide(login.NewloginService)

	//pigeon
	container.Provide(orm.NewPigeonRepo)
	container.Provide(pigeon.NewpigeonService)

	//ptstatus
	container.Provide(orm.NewPtstatusRepo)
	container.Provide(ptstatus.NewpigeonService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
