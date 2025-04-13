package internal

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dmytro-kucherenko/smartner-api-gateway/internal/authorizer"
)

func Init() {
	lambda.Start(authorizer.Handle)
}
