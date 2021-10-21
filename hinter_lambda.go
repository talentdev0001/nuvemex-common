package part

import (
	"encoding/json"
	"github.com/Montrealist-cPunto/commons/log"
	"github.com/Montrealist-cPunto/goseanto"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

type HinterLambda struct {
	Service *goseanto.Hinter
	Logger  *log.Logger
}

type HinterRequestParameters struct {
	PartNum string `json:"partNum"`
	Field   string `json:"field"`
}

func (h *HinterLambda) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer h.Logger.Await(LogFlushWait)

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

	params := &HinterRequestParameters{}
	js, _ := json.Marshal(request.QueryStringParameters)
	err := json.Unmarshal(js, params)
	if err != nil {
		h.Logger.Debug().Err(err).Send()
		return response, nil
	}

	keyword := params.PartNum
	field := params.Field
	limit, _ := strconv.Atoi(request.QueryStringParameters["limit"])

	if keyword == "" {
		return response, nil
	}

	if field == "" {
		field = "partNum.raw"
	}

	results := h.Service.Get(&goseanto.HinterOptions{
		Field:      field,
		PartNumber: keyword,
		Limit:      limit,
	})

	encoded, err := json.Marshal(results)
	if err != nil {
		h.Logger.Err(err)
		return response, nil
	}
	response.Body = string(encoded)

	return response, nil
}
