//
// -- Netlify Functions
//

package submit

import (
	"net/http"
	"github.com/aws/aws-lambda-go/events"
)

type response struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Login  string `json:"login"`
	Status string `json:"status"`
}

func Handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	body := response{
		ID: "ID",
		UserID: "userID",
		Login: "Login",
		Status: "Status",
	}
	return apiResponse(http.StatusOK, body)
}


