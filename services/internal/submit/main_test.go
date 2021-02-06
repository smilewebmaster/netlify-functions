//
// -- Submit Handler
//

package submit

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestApiHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ApiHandler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApiHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiResponse(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApiResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApiResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
