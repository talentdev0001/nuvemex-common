package main

import (
	"github.com/Montrealist-cPunto/commons/elasticsearch"
	part "github.com/Montrealist-cPunto/gos-part"
	"github.com/Montrealist-cPunto/goseanto"
	"strings"
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
