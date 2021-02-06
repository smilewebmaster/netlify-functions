//
// -- Submit Handler
//

package submit

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  Result
}

type Result struct {
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Data      string
}

func toJson(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}

func ApiHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	// if no name is provided in the HTTP request body, throw an error
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}
	// struct
	responseData := Response{
		Status:  200,
		Message: "",
		Result: Result{
			User:      "bbg0x",
			CreatedAt: time.Time{},
			Data:      crawler(),
		},
	}
	// response
	resp := events.APIGatewayProxyResponse{}
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	resp.Body = toJson(responseData)
	resp.StatusCode = 200
	return resp, nil
}
