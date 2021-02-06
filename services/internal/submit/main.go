//
// -- Submit Handler
//

package submit

import (
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func ApiHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	// if no name is provided in the HTTP request body, throw an error
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}
	// response
	resp := events.APIGatewayProxyResponse{}
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	resp.Body = crawler()
	resp.StatusCode = 200
	return resp, nil
}
