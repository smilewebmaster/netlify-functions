//
// -- Submit Handler
//

package submit

import (
    "log"
    "errors"
    "github.com/aws/aws-lambda-go/events"
)

// type response struct {
//     ID     string `json:"id"`
//     UserID string `json:"user_id"`
//     SiteUrl string `json:"site_url"`
// }

// // BodyRequest is our self-made struct to process JSON request from Client
// type request struct {
//     RequestName string `json:"name"`
// }

// var (
//     // ErrNameNotProvided is thrown when a name is not provided
//     ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
// )

// func ApiHandler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//     if(req.HTTPMethod != "POST") {
//         return events.APIGatewayProxyResponse{}, ErrNameNotProvided
//     }
//     body := response{ ID: "id", UserID: "user_id", SiteUrl: "http://example.com" }
//     return apiResponse(http.StatusOK, body)
// }


var (
    ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func ApiHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // stdout and stderr are sent to AWS CloudWatch Logs
    log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
    // If no name is provided in the HTTP request body, throw an error
    if request.HTTPMethod != "POST" {
        return events.APIGatewayProxyResponse{}, ErrNameNotProvided
    }
    resp := events.APIGatewayProxyResponse{}
    resp.Headers = map[string]string{"Content-Type": "application/json"}
    resp.Body = request.Body.SiteUrl
    resp.StatusCode = 200
    return resp, nil
}

//
// -- END
//

// app := &response{
//     Name:  "build-lambda-zip",
//     Usage: "Put an executable and supplemental files into a zip file that works with AWS Lambda.",
//     Flags: []cli.Flag{
//         &cli.StringFlag{
//             Name:    "output",
//             Aliases: []string{"o"},
//             Value:   "",
//             Usage:   "output file path for the zip. Defaults to the first input file name.",
//         },
//     },
//     Action: func(c *cli.Context) error {
//         if !c.Args().Present() {
//             return errors.New("no input provided")
//         }

//         inputExe := c.Args().First()
//         outputZip := c.String("output")
//         if outputZip == "" {
//             outputZip = fmt.Sprintf("%s.zip", filepath.Base(inputExe))
//         }

//         if err := compressExeAndArgs(outputZip, inputExe, c.Args().Tail()); err != nil {
//             return fmt.Errorf("failed to compress file: %v", err)
//         }
//         log.Print("wrote " + outputZip)
//         return nil
//     },
// }