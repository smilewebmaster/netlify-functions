//
// -- Site Crawler
//

package submit

import (
	"reflect"
	"testing"
)

func Test_crawler(t *testing.T) {
	type args struct {
		siteUrl string
	}
	tests := []struct {
		name string
		args args
		want PageData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crawler(tt.args.siteUrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crawler() = %v, want %v", got, tt.want)
			}
		})
	}
}
