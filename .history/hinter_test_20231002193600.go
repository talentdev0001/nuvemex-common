package part_test

import (
	"strings"
	"testing"
	"time"

	"github.com/nuvemex/commons/elasticsearch"
	"github.com/nuvemex/commons/math"
	part "github.com/nuvemex/gos-part"
	"github.com/nuvemex/goseanto"
	"github.com/stretchr/testify/assert"
)

func TestHinter(t *testing.T) {
	cfg := part.MustConfig()
	index := elasticsearch.MustIndex(cfg)
	schema := goseanto.RootPath + "/resources/test/elastic/test-part-mapping.json"

	t.Run("Checks hint results received and sent to crawl queue", func(t *testing.T) {
		partNumber := "bav99"
		service := part.MustHinterService(cfg)
		elastic := goseanto.MustElasticSearch(cfg)
		persister := goseanto.MustPersister(cfg)

		indexName := "test-index-" + strings.ToLower(math.RandomString(6))
		persister.Config.Index.Name = indexName
		elastic.Config.Index.Name = indexName
		_, err := index.Create(indexName, schema)
		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			_, _ = index.Delete(indexName)
		}()
		service.ElasticService = elastic

		doc1 := &goseanto.Record{
			PartNum:  "mav99",
			Supplier: "arrow",
		}
		doc2 := &goseanto.Record{
			PartNum:  "bav98",
			Supplier: "arrow",
		}
		doc3 := &goseanto.Record{
			PartNum:  "bav99",
			Supplier: "arrow",
		}
		doc31 := &goseanto.Record{
			PartNum:      "bav99",
			Manufacturer: "xxx",
			Supplier:     "arrow",
		}
		doc4 := &goseanto.Record{
			PartNum:  "abc99",
			Supplier: "arrow",
		}
		doc5 := &goseanto.Record{
			PartNum:  "bazzz",
			Supplier: "arrow",
		}
		persister.Save([]*goseanto.Record{
			doc1,
			doc2,
			doc3,
			doc31,
			doc4,
			doc5,
		})

		// give some time to update index
		time.Sleep(900 * time.Millisecond)

		expected := []*goseanto.HintResult{
			{Value: "bav99", Total: 2},
		}

		got := service.Get(&goseanto.HinterOptions{
			Field:      "partNum.raw",
			PartNumber: partNumber,
			Limit:      10,
		})
		assert.Equal(t, expected, got)
	})
}
