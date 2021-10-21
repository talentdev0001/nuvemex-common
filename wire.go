//+build wireinject

package part

import (
	"errors"
	"fmt"
	"github.com/Montrealist-cPunto/commons/config"
	"github.com/Montrealist-cPunto/commons/log"
	"github.com/Montrealist-cPunto/goseanto"
	"github.com/google/wire"
	"os"
	"sync"
)

var onceLibConfig sync.Once
var libConfig *config.Config

func ProvideLibConfig(appConfig *AppConfig) *config.Config {
	if libConfig == nil {
		panic(errors.New("libConfig not inited"))
	}
	return libConfig
}

var onceAppConfig sync.Once
var appConfig *AppConfig

func ProvideAppConfig(dir string) *AppConfig {
	onceAppConfig.Do(func() {
		cfg := config.LoadFromDirectory(dir)
		appConfig = &AppConfig{
			Config: cfg,
		}
	})
	onceLibConfig.Do(func() {
		libConfig = config.LoadFromDirectory(dir,
			"goseanto.yml",
			fmt.Sprintf("goseanto-%s.yml", os.Getenv("app_env")))
		// overwrite parent config
		_ = libConfig.MergeAt(appConfig.Config.Koanf, "log")
	})

	return appConfig
}

//func provideSearchConfig(cfg *AppConfig) *SearchConfig {
//	serviceConfig := &SearchConfig{}
//	err := cfg.Unmarshal("app.search", serviceConfig)
//	if err != nil {
//		panic(err)
//	}
//	return serviceConfig
//}

func MustSearchLambda(appConfig *AppConfig) *SearchLambda {
	panic(wire.Build(
		ProvideLibConfig,
		goseanto.MustSearchService,
		log.MustLogger,
		wire.Struct(new(SearchLambda), "*")))
}

func MustHinterLambda(appConfig *AppConfig) *HinterLambda {
	panic(wire.Build(
		ProvideLibConfig,
		goseanto.MustHinterService,
		log.MustLogger,
		wire.Struct(new(HinterLambda), "*")))
}

func MustDetailsLambda(appConfig *AppConfig) *DetailsLambda {
	panic(wire.Build(
		ProvideLibConfig,
		goseanto.MustSearchService,
		log.MustLogger,
		wire.Struct(new(DetailsLambda), "*")))
}
