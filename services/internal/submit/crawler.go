//
// -- Site Crawler
//

package submit

import (
	"log"
	"net/http"

	"github.com/foolin/pagser"
)

type PageData struct {
	Title     string `json:"title" pagser:"title"`
	Canonical string `json:"canonical" pagser:"link[rel=canonical]->attr(href)"`
	Meta      struct {
		Description string   `json:"description" pagser:"meta[name='description']->attr(content)"`
		Keywords    []string `json:"keywords" pagser:"meta[name='keywords']->attrSplit(content)"`
	} `pagser:"head"`
	OpenGraph struct {
		Locale      string `json:"locale" pagser:"meta[property='og:locale']->attr(content)"`
		Type        string `json:"type" pagser:"meta[property='og:type']->attr(content)"`
		Title       string `json:"title" pagser:"meta[property='og:title']->attr(content)"`
		Description string `json:"description" pagser:"meta[property='og:description']->attr(content)"`
		Url         string `json:"url" pagser:"meta[property='og:url']->attr(content)"`
		SiteName    string `json:"site_name" pagser:"meta[property='og:site_name']->attr(content)"`
		Image       string `json:"image" pagser:"meta[property='og:image']->attr(content)"`
	} `pagser:"head"`
}

func isFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var data PageData

func crawler(siteUrl string) PageData {
	response, err := http.Get(siteUrl)
	isFatal(err)
	defer response.Body.Close()
	html := pagser.New()
	err = html.ParseReader(&data, response.Body) // parse html data
	// check error
	isFatal(err)
	return data // print data
}

//
// -- END
//
