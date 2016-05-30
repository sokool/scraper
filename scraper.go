package main

import (
	"github.com/sokool/scraper/requestor"
	"github.com/sokool/scraper/crawler"
)

func main() {

	//sites := []string {
	//	"http://google.pl",
	//	"http://onet.pl",
	//	"http://wp.pl",
	//	"http://polsat.pl",
	//	"http://tvn.pl",
	//	"http://facebook.pl",
	//	"http://fakty.interia.pl/polska",
	//	"http://mint-soft.pl",
	//	"http://google.de",
	//	"http://www.wroclaw.apodatkowa.gov.pl/izba-skarbowa-we-wroclawiu;jsessionid=C3FE209C723806A6F34547C5EF3C5C34",
	//}
	//
	//showTitle := func(page *requestor.Page) {
	//	d := page.Document();
	//	fmt.Printf("%s://%s [%s] %s\n", d.Url.Scheme, d.Url.Host, page.LoadingTime(), d.Find("title").Text())
	//}
	//
	//request := requestor.New()
	//
	//for _, url := range sites {
	//	request.Do(url, showTitle)
	//}
	//request.WaitForAll()
	//
	//finish()

	//proc := make(chan *Ta)

	//for _, c := range r.configs {
	//	url := newUrl(c.Domain, c.URI, func(p *page) {
	//		o := &Ta{page: p, conf: c}
	//		r.list(o)
	//
	//	})
	//
	//	r.request(url)
	//
	//}
	//r.finish()

	olx := &crawler.Configuration{
		URL: "http://olx.pl/oferty/",
		Next: "a[class*=pageNextPrev]",
		Object: "a[class*='thumb']",
		Template: crawler.Template{
			"id" : "span[class='nowrap marginright'] span",
			"name" : "div[class*=offerheadinner] h1",
			//"description" : "div[id*=textContent] p",
			"createdAt" : "span[class='pdingleft10 brlefte5']",
			"price" : crawler.Query(func(p *requestor.Page) string {
				o := p.Document().Find("div[class*='pricelabel'] strong").First().Text()
				return o
			}),
		},
	}
	otomoto := &crawler.Configuration{
		URL: "http://otomoto.pl/osobowe",
		Next: "[class='next abs'] a",
		Object: "a[class*='img-cover']",
		Template: crawler.Template{
			"name" : "header[class*=om-offer-title] div[class=row] h1",
			"cena" : "div[class=price-cell] span[class=om-price]",
		},
	}
	//
	homegate := &crawler.Configuration{
		URL: "http://www.homegate.ch/buy/apartment/matching-list?lastMap=ctn_gr&ab=G000000000000000000000000000000000000000000000000000000000000000000000000000000120000000008E00200C082640001000000000000&ep=%d",
		Next: "a[rel=next]",
		Object: "a[class='detail-page-link box-row--link']",
		Template: crawler.Template{
			"id" : "div[class=nr] span",
			"name" : "h1.title",
			//"description" : "div[class='description-content']",
		},

	}

	crawler.
	New().
	Add(olx).
	Add(otomoto).
	Add(homegate).
	Run()

}
