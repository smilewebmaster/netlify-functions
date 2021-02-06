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
	Title     string `pagser:"title"`
	Canonical string `pagser:"link[rel=canonical]->attr(href)"`
	Meta      struct {
		Description string   `pagser:"meta[name='description']->attr(content)"`
		Keywords    []string `pagser:"meta[name='keywords']->attrSplit(content)"`
	} `pagser:"head"`
	OpenGraph struct {
		Locale      string `pagser:"meta[property='og:locale']->attr(content)"`
		Type        string `pagser:"meta[property='og:type']->attr(content)"`
		Title       string `pagser:"meta[property='og:title']->attr(content)"`
		Description string `pagser:"meta[property='og:description']->attr(content)"`
		Url         string `pagser:"meta[property='og:url']->attr(content)"`
		SiteName    string `pagser:"meta[property='og:site_name']->attr(content)"`
		Image       string `pagser:"meta[property='og:image']->attr(content)"`
	} `pagser:"head"`
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func crawler() PageData {
	response, err := http.Get("https://medium.com/swlh/write-a-custom-reusable-hook-usefetch-1443d8d4e1e1")
	fatal(err)
	defer response.Body.Close()
	page := pagser.New()
	var data PageData
	// parse html data
	err = page.ParseReader(&data, response.Body)
	// check error
	fatal(err)
	// print data
	return data
}

//
// -- END
//
