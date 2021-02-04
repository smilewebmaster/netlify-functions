
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bbg/netlify-functions/services/internal/hello"
)


func main() {
        lambda.Start(hello.HandleLambdaEvent)
}