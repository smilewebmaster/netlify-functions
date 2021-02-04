//
// -- Netlify Functions
//

package main

import (
	"encoding/json"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Login  string `json:"login"`
	Status string `json:"status"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := response{
		ID: "ID",
		UserID: "userID",
		Login: "Login",
		Status: "Status",
	}
	response, err := json.Marshal(&body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: err.Error(), 
			StatusCode: 404,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body: string(response),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
