package part

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nuvemex/commons/log"
	"github.com/nuvemex/goseanto"
)

type DetailsLambda struct {
	Service *goseanto.SearchService
	Logger  *log.Logger
}

func (d *DetailsLambda) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer d.Logger.Await(LogFlushWait)
	response := events.APIGatewayProxyResponse{
		Body: "{}",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: 200,
	}

	partID := request.QueryStringParameters["partId"]
	if len(partID) < 32 {
		return response, nil
	}

	d.Logger.Debug().Msg(fmt.Sprintf("Get: %s", partID))

	result, err := d.Service.GetByID(partID)

	// gracefully indicate there's nothing found
	if err != nil {
		d.Logger.Debug().Msgf("failed to get pat by ID %s", partID)
		response.StatusCode = http.StatusNoContent
		return response, nil
	}

	encoded, _ := json.Marshal(result)
	response.Body = string(encoded)

	return response, nil
}
