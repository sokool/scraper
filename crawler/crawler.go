package crawler

import (
	. "github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/sokool/scraper/requestor"
	"net/url"
)

type Template map[string]Query

type Query func(*requestor.Page) interface{}

type Configuration struct {
	URL      string
	Next     string
	Object   string
	Template Template
	receiver func(map[string]interface{})
};

type Crawler struct {
	configs []*Configuration
	request *requestor.Request
	finish  func()
	counter int
}

func (r *Crawler) Add(config *Configuration, fn func(record map[string]interface{})) *Crawler {
	config.receiver = fn
	r.configs = append(r.configs, config)

	return r
}

func (r *Crawler) visitObject(p *requestor.Page, c *Configuration) {
	output := make(map[string]interface{})
	output["url"] = p.Document().Url.String()
	for name, selector := range c.Template {
		output[name] = selector(p)
	}
	r.counter++

	c.receiver(output)
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

func Each(in, key, value string) Query {
	return Query(func(page *requestor.Page) interface{} {
		out := make(map[string]string)
		page.Document().Find(in).Each(func(i int, item *Selection) {
			left := item.Find(key).Text()
			if left != "" {
				out[item.Find(key).Text()] = item.Find(value).Text()
			}
		})

		return out
	})
}

func First(in string) Query {
	return Query(func(page *requestor.Page) interface{} {
		return page.Document().Find(in).First().Text()
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