//
// -- Api Response Function
//

package submit

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{}
	resp.Headers = map[string]string{"Content-Type": "application/json"}
	resp.StatusCode = status
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

//
// -- END
//
