package part

import (
	"strings"

	"github.com/nuvemex/commons/log"
	"github.com/nuvemex/commons/queue"
	"github.com/nuvemex/goseanto"
)

type Hinter struct {
	ElasticService *goseanto.ElasticSearch
	Queue          *queue.Queue
	Suppliers      []*goseanto.Supplier
	Logger         *log.Logger
}

func (h *Hinter) Get(options *HinterOptions) []*HintResult {
	suppliers := options.Suppliers
	if len(suppliers) == 0 {
		suppliers = h.getAvailableSuppliers()
	}
	options.Suppliers = suppliers

	res := h.ElasticService.GetHints(options)
	ln := len(res)
	if ln == 0 {
		h.Logger.Debug().Msg("no hints returned from es.")
		return res
	}

	return res
}

func (h *Hinter) getAvailableSuppliers() []string {
	res := make([]string, 0)

	for _, supplier := range h.Suppliers {
		res = append(res, strings.ToLower(supplier.Name))
	}

	return res
}
