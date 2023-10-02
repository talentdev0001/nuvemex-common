package part

import (
	"encoding/json"
	"github.com/Montrealist-cPunto/commons/log"
	"github.com/Montrealist-cPunto/goseanto"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
	"strings"
)

type SearchLambda struct {
	Service *goseanto.SearchService
	Logger  *log.Logger
}

func (s *SearchLambda) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer s.Logger.Await(LogFlushWait)

	response := events.APIGatewayProxyResponse{
		Body: "[]",
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		StatusCode: 200,
	}

	keyword := request.QueryStringParameters["partNum"]
	if keyword == "" {
		return response, nil
	}

	sort := request.QueryStringParameters["sort"]
	sorts := strings.Split(sort, ",")

	js2, _ := json.Marshal(sorts)
	s.Logger.Debug().Msg("sorts: " + string(js2))

	limit, _ := strconv.Atoi(request.QueryStringParameters["limit"])
	offset, _ := strconv.Atoi(request.QueryStringParameters["offset"])
	inStock, _ := strconv.Atoi(request.QueryStringParameters["inStock"])

	suppliersQuery := request.QueryStringParameters["supplier"]
	suppliers := make([]string, 0)
	if suppliersQuery != "" {
		suppliers = strings.Split(suppliersQuery, ",")
	}
	if len(suppliers) > 100 {
		suppliers = suppliers[:100]
	}

	results := s.Service.Search(&goseanto.SearchOptions{
		PartNumber: keyword,
		Sort:       sorts,
		Pagination: &goseanto.SearchPagination{
			Offset: offset,
			Limit:  limit,
		},
		RequiredStockAmount: inStock,
		Suppliers:           suppliers,
	})

	encoded, _ := json.Marshal(results)
	response.Body = string(encoded)

	return response, nil
}
