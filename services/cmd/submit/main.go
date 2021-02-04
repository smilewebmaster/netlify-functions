package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbg/netlify-functions/services/internal/submit"
)

func main() {
	lambda.Start(submit.Handler)
}