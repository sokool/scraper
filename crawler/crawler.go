package crawler

import (
	. "github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/sokool/scraper/requestor"
	"net/url"
	"strings"
	"encoding/json"
	"os"
)

type Template map[string]interface{}

type Query func(*requestor.Page) interface{}

type Configuration struct {
	URL      string
	Next     string
	Object   string
	Template Template
	ExportFile string
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
func (r *Crawler) selectorValue(page *requestor.Page, input interface{}) interface{} {
	switch o := input.(type) {
	case Query:
		return o(page)
	case string:
		return strings.TrimSpace(page.Document().Find(o).Text())
	default:
		return ""
	}
}

func (r *Crawler) visitObject(p *requestor.Page, c *Configuration) {
	o := Template{}
	o["url"] = p.Document().Url.String()
	for name, selector := range c.Template {
		o[name] = r.selectorValue(p, selector)
	}
	r.counter++
	//fmt.Printf("[%d]. %s\n", r.counter, o)
	js, _ := json.Marshal(o)
	os.Stdout.Write(js)
}

func (r *Crawler) visitRows(c *Configuration, p *requestor.Page) {
	p.Document().Find(c.Object).Each(func(i int, item *Selection) {
		uri, _ := item.Attr("href")
		r.request.Do(fqdn(uri, c.URL), func(page *requestor.Page) {
			r.visitObject(page, c)
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

func EachText(in, key, value string) Query {
	return Query(func(page *requestor.Page) interface{} {
		out := make(map[string]string)
		page.Document().Find(in).Each(func(i int, item *Selection) {
			out[item.Find(key).Text()] = item.Find(value).Text()
		})

		return out
	})
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

	return fmt.Sprintf("%s://%s%s", b.Scheme, b.Host, a)

}