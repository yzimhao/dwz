package html

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type Html struct {
	url         string
	title       string
	keyword     string
	description string

	source string
}

func NewHtml(url string) *Html {
	h := &Html{}
	h.url = url
	h.request()
	return h
}

func (h *Html) request() {
	// Request the HTML page.
	res, err := http.Get(h.url)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	// Find the review items
	h.source = doc.Text()
	h.title = doc.Find("title").Text()
	h.keyword = doc.Find("keyword").Text()
	h.description = doc.Find("description").Text()
}

func (h *Html) Url() string {
	return h.url
}

func (h *Html) Title() string {
	return h.title
}
func (h *Html) Keyword() string {
	return h.keyword
}
func (h *Html) Description() string {
	return h.description
}
func (h *Html) Source() string {
	return h.source
}
