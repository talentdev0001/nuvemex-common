package main

import (
	ddlambda "github.com/DataDog/datadog-lambda-go"
	"github.com/aws/aws-lambda-go/lambda"
	part "github.com/nuvemex/gos-part"
)

func main() {
	handler := part.MustDetailsLambda(part.MustConfig())

	defer handler.Logger.Await(part.LogFlushWait)
	lambda.Start(ddlambda.WrapFunction(handler.Handle, nil))
}
