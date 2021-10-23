// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package part

import (
	"fmt"
	"github.com/Montrealist-cPunto/commons/config"
	"github.com/Montrealist-cPunto/commons/log"
	"github.com/Montrealist-cPunto/goseanto"
	"os"
	"sync"
)

// Injectors from wire.go:

func MustSearchLambda(appConfig2 *config.Config) *SearchLambda {
	searchService := goseanto.MustSearchService(appConfig2)
	logger := log.MustLogger(appConfig2)
	searchLambda := &SearchLambda{
		Service: searchService,
		Logger:  logger,
	}
	return searchLambda
}

func MustHinterLambda(appConfig2 *config.Config) *HinterLambda {
	hinter := goseanto.MustHinterService(appConfig2)
	logger := log.MustLogger(appConfig2)
	hinterLambda := &HinterLambda{
		Service: hinter,
		Logger:  logger,
	}
	return hinterLambda
}

func MustDetailsLambda(appConfig2 *config.Config) *DetailsLambda {
	searchService := goseanto.MustSearchService(appConfig2)
	logger := log.MustLogger(appConfig2)
	detailsLambda := &DetailsLambda{
		Service: searchService,
		Logger:  logger,
	}
	return detailsLambda
}

// wire.go:

var onceAppConfig sync.Once

var appConfig *config.Config

func MustConfig() *config.Config {
	onceAppConfig.Do(func() {
		appConfig = config.LoadFromDirectory("./resources/config",
			"goseanto.yml", fmt.Sprintf("goseanto-%s.yml", os.Getenv("app_env")), "config.yml", fmt.Sprintf("%s.yml", os.Getenv("app_env")),
		)
	})

	return appConfig
}
