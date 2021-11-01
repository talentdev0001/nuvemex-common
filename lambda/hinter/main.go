package main

import (
	ddlambda "github.com/DataDog/datadog-lambda-go"
	part "github.com/Montrealist-cPunto/gos-part"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := part.MustHinterLambda(part.MustConfig())

	defer handler.Logger.Await(part.LogFlushWait)
	lambda.Start(ddlambda.WrapFunction(handler.Handle, nil))
}
