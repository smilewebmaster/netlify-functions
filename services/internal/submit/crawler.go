//
// -- Site Crawler
//

package submit

import (
	"encoding/json"
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

func toJson(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func crawler() string {
	resp, err := http.Get("https://medium.com/swlh/write-a-custom-reusable-hook-usefetch-1443d8d4e1e1")
	fatal(err)
	defer resp.Body.Close()
	p := pagser.New()
	var data PageData
	// parse html data
	err = p.ParseReader(&data, resp.Body)
	// check error
	fatal(err)
	// print data
	return toJson(data)
}

//
// -- END
//
