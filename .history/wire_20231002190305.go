//go:build wireinject
// +build wireinject

package part

import (
	"fmt"
	"os"
	"sync"

	"github.com/google/wire"
	"github.com/nuvemex/commons/config"
	"github.com/nuvemex/commons/log"
	"github.com/nuvemex/commons/queue"
	"github.com/nuvemex/goseanto"
)

var onceAppConfig sync.Once
var appConfig *config.Config

func MustConfig() *config.Config {
	onceAppConfig.Do(func() {
		appConfig = config.LoadFromDirectory("./resources/config",
			"goseanto.yml",
			fmt.Sprintf("goseanto-%s.yml", os.Getenv("app_env")),
			"config.yml",
			fmt.Sprintf("%s.yml", os.Getenv("app_env")),
		)
	})

	return appConfig
}

var onceHinterService sync.Once
var hinterService *Hinter

func provideHinterService(cfg *config.Config) *Hinter {
	panic(wire.Build(
		goseanto.MustElasticSearch,
		log.MustLogger,
		queue.MustQueue,
		goseanto.ProviderSuppliers,
		wire.Struct(new(Hinter), "*")))
}

func MustHinterService(cfg *config.Config) *Hinter {
	onceHinterService.Do(func() {
		hinterService = provideHinterService(cfg)
	})

	return hinterService
}

func MustSearchLambda(appConfig *config.Config) *SearchLambda {
	panic(wire.Build(
		goseanto.MustSearchService,
		log.MustLogger,
		wire.Struct(new(SearchLambda), "*")))
}

func MustHinterLambda(appConfig *config.Config) *HinterLambda {
	panic(wire.Build(
		MustHinterService,
		log.MustLogger,
		wire.Struct(new(HinterLambda), "*")))
}

func MustDetailsLambda(appConfig *config.Config) *DetailsLambda {
	panic(wire.Build(
		goseanto.MustSearchService,
		log.MustLogger,
		wire.Struct(new(DetailsLambda), "*")))
}
