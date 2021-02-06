//
// -- Site Crawler
//

package submit

import "testing"

func Test_crawler(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crawler(); got != tt.want {
				t.Errorf("crawler() = %v, want %v", got, tt.want)
			}
		})
	}
}
