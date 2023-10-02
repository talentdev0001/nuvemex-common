package main

import (
	"strings"

	part "github.com/Montrealist-cPunto/gos-part"
	"github.com/nuvemex/commons/elasticsearch"
	"github.com/nuvemex/goseanto"
)

func main() {
	cfg := part.MustConfig()

	serviceConfig := elasticsearch.MustConfig(cfg)

	service := elasticsearch.MustIndex(cfg)
	_, err := service.Create(
		serviceConfig.Index.Name,
		strings.Replace(serviceConfig.Index.Schema, "{RootPath}", goseanto.RootPath, 1))

	if err != nil {
		println("elasticsearch create error:", err.Error())
		return
	}

	println("Installed elasticsearch mapping")
}
