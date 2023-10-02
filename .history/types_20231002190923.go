package part

import "github.com/nuvemex/goseanto"

type SearchConfig struct {
	CacheDuration string
	CrawlPool     string
	URL           string
}

type HinterOptions = goseanto.HinterOptions
type HintResult = goseanto.HintResult
