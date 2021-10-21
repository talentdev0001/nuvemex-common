package part

import (
	"github.com/Montrealist-cPunto/commons/config"
)

type AppConfig struct {
	*config.Config
}

type SearchConfig struct {
	CacheDuration string
	CrawlPool     string
	URL           string
}
