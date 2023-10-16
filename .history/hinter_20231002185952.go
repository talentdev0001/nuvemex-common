package part

import (
	"github.com/Montrealist-cPunto/commons/log"
	"github.com/Montrealist-cPunto/commons/queue"
	"github.com/Montrealist-cPunto/goseanto"
	"strings"
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
