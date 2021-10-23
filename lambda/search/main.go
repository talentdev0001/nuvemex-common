package main

import (
	part "github.com/Montrealist-cPunto/gos-part"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	handler := part.MustSearchLambda(part.MustConfig())

	defer handler.Logger.Await(part.LogFlushWait)
	lambda.Start(handler.Handle)
}
