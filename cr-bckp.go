package xx

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
)

type Meta struct {
	Title       string
	Description string
}

func fatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func xpathRead(path string, reader io.Reader) string {
	xmlroot, xmlerror := xmlpath.ParseHTML(reader)
	fatalError(xmlerror)
	xmlcontent := xmlpath.MustCompile(path)
	if value, ok := xmlcontent.String(xmlroot); ok {
		return value
	}
	return "Not Found"
}

func main() {
	resp, err := http.Get("https://kananrahimov.com/")
	fatalError(err)
	defer resp.Body.Close()
	//
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fatalError(err)
	//
	reader := strings.NewReader(string(bodyBytes))
	root, err := html.Parse(reader)
	fatalError(err)
	//
	var b bytes.Buffer
	html.Render(&b, root)
	fixedHtml := b.String()
	reader = strings.NewReader(fixedHtml)

	// title := xmlpath.MustCompile(`//html/head/title`)
	// description := xmlpath.MustCompile(`//head/meta[@name='description']/@content`)

	// fmt.Println(xpathRead(`//html/head/title`, reader))
	// fmt.Println(xpathRead(`//head/meta[@name='description']/@content`, reader))

	// fmt.Println(title.String(xmlroot))
	// fmt.Println(description.String(xmlroot))

	fmt.Println(xpathRead(`//html/head/title`, reader))
	fmt.Println(xpathRead(`//head/meta[@name='description']/@content`, reader))
}
