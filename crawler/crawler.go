package crawler

import (
	. "github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/sokool/scraper/requestor"
	"net/url"
	"strings"
)

type Template map[string]interface{}

type Query func(*requestor.Page) string

type Configuration struct {
	URL      string
	Next     string
	Object   string
	Template Template
};

type Crawler struct {
	configs []*Configuration
	request *requestor.Request
	finish  func()
	counter int
}

func (r *Crawler) Add(config *Configuration) *Crawler {
	r.configs = append(r.configs, config)

	return r
}
func (r *Crawler) selectorValue(page *requestor.Page, input interface{}) string {
	switch o := input.(type) {
	case Query:
		return o(page)
	case string:
		return strings.TrimSpace(page.Document().Find(o).Text())
	default:
		return ""
	}
}

func (r *Crawler) visitObject(p *requestor.Page, t Template) {
	o := Template{}
	o["url"] = p.Document().Url.String()
	for name, selector := range t {
		o[name] = r.selectorValue(p, selector)
	}
	r.counter++
	fmt.Printf("[%d]. %s\n", r.counter, o)
}

func (r *Crawler) visitRows(c *Configuration, p *requestor.Page) {
	p.Document().Find(c.Object).Each(func(i int, item *Selection) {
		uri, _ := item.Attr("href")
		r.request.Do(fqdn(uri, c.URL), func(page *requestor.Page) {
			r.visitObject(page, c.Template)
		})
	});

}

func (r *Crawler) visitResult(url string, c *Configuration) {
	r.request.Do(fqdn(url, c.URL), func(p *requestor.Page) {
		nextUri, ok := p.Document().Find(c.Next).Attr("href")
		if ok {
			r.visitResult(fqdn(nextUri, c.URL), c)
		}
		r.visitRows(c, p)
	})

}

func (r *Crawler) Run() {
	for _, conf := range r.configs {
		r.visitResult(conf.URL, conf)
	}
	r.request.WaitForAll()

}

func New() *Crawler {
	c := &Crawler{
		configs: make([]*Configuration, 0),
		request: requestor.New(),
		counter: 0,
	}
	return c
}

func fqdn(urlA, urlB string) string {
	a, _ := url.Parse(urlA)
	b, _ := url.Parse(urlB)

	if a.Host != "" {
		return urlA
	}

	return fmt.Sprintf("%s://%s/%s", b.Scheme, b.Host, a)

}