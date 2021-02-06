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
	Result  struct {
		User      string    `json:"user"`
		CreatedAt time.Time `json:"created_at"`
		Data      PageData  `json:"data"`
	} `json:"result"`
}

func toJson(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}

func ApiResponse() Response {
	// response struct
	responseData := Response{}
	responseData.Status = 200
	responseData.Message = ""
	responseData.Result.User = "bbg0x"
	responseData.Result.CreatedAt = time.Now()
	responseData.Result.Data = crawler("https://medium.com/swlh/write-a-custom-reusable-hook-usefetch-1443d8d4e1e1")
	// end
	return responseData
}

func ApiHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	// if no name is provided in the HTTP request body, throw an error
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}
	// response
	response := events.APIGatewayProxyResponse{}
	response.Headers = map[string]string{"Content-Type": "application/json"}
	response.Body = toJson(ApiResponse())
	response.StatusCode = 200
	return response, nil
}
