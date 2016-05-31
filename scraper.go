package main

import (
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

	//olx := &crawler.Configuration{
	//	URL: "http://olx.pl/oferty/",
	//	Next: "a[class*=pageNextPrev]",
	//	Object: "a[class*='thumb']",
	//	Template: crawler.Template{
	//		"id" : "span[class='nowrap marginright'] span",
	//		"name" : "div[class*=offerheadinner] h1",
	//		//"description" : "div[id*=textContent] p",
	//		"createdAt" : "span[class='pdingleft10 brlefte5']",
	//		//"price" : crawler.Query(func(p *requestor.Page) string {
	//		//	o := p.Document().Find("div[class*='pricelabel'] strong").First().Text()
	//		//	return o
	//		//}),
	//	},
	//}
	//immoscout := &crawler.Configuration{
	//	URL: "http://www.immobilienscout24.de/Suche/S-2/Wohnung-Kauf?enteredFrom=result_list",
	//	Next: "a[data-is24-qa='paging_bottom_next']",
	//	Object: "a[class='result-list-entry__brand-title-container']",
	//	Template: crawler.Template{
	//		"name" : "#expose-title",
	//		"price": "div[class*='is24qa-kaufpreis is24-value']",
	//		"type" : "dd[class*='is24qa-wohnungstyp']",
	//		"living_space" : "div[class*='is24qa-wohnflaeche-ca is24-value']",
	//		"rooms": "div[class*='is24qa-zi is24-value']",
	//		//"description" : "pre[class=is24qa-objektbeschreibung]",
	//	},
	//}
	//
	//otomoto := &crawler.Configuration{
	//	URL: "http://otomoto.pl/osobowe",
	//	Next: "[class='next abs'] a",
	//	Object: "a[class*='img-cover']",
	//	Template: crawler.Template{
	//		"name" : "header[class*=om-offer-title] div[class=row] h1",
	//		"cena" : "div[class=price-cell] span[class=om-price]",
	//		"ficzery" : crawler.EachText(
	//			"ul[class*='params-list'] li",
	//			"small",
	//			"span",
	//		),
	//	},
	//}
	//

	homegate := &crawler.Configuration{
		URL: "http://www.homegate.ch/buy/apartment/matching-list?lastMap=ctn_gr&ab=G000000000000000000000000000000000000000000000000000000000000000000000000000000120000000008E00200C082640001000000000000",
		Next: "a[rel=next]",
		Object: "a[class='detail-page-link box-row--link']",
		ExportFile: "homegate.json",
		Template: crawler.Template{
			"id" : "div[class=nr] span",
			"name" : "h1.title",
			"price" : "span[itemprop='price']",
			//"address": "div[class*=detail-address]",
			//"type" : crawler.First{"ul[class='list--plain list--flat list--spaced-double text--small'] span[class*='float-right']"},
			"features" : crawler.EachText(
				"ul[class='list--plain list--flat list--spaced-double text--small'] li",
				"span[class='text--small']",
				"span[class*='float-right']",
			),
			//"description" : "div[class='description-content']",
		},

	}

	crawler.
	New().
	//Add(olx).
	//Add(immoscout).
	//Add(otomoto).
	Add(homegate).
	Run()

}
