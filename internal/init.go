package internal

import (
	"github.com/Dmytro-Kucherenko/smartner-api-gateway/internal/authorizer"
	"github.com/aws/aws-lambda-go/lambda"
)

func Init() {
	lambda.Start(authorizer.Handle)
}
