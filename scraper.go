//package main
//
//import (
//	. "github.com/sokool/scraper/crawler"
//	"encoding/json"
//	"fmt"
//	"github.com/sokool/scraper/tag"
//)
//
//func main() {
//
//	//olx := &crawler.Configuration{
//	//	URL: "http://olx.pl/oferty/",
//	//	Next: "a[class*=pageNextPrev]",
//	//	Object: "a[class*='thumb']",
//	//	Template: crawler.Template{
//	//		"id" : "span[class='nowrap marginright'] span",
//	//		"name" : "div[class*=offerheadinner] h1",
//	//		//"description" : "div[id*=textContent] p",
//	//		"createdAt" : "span[class='pdingleft10 brlefte5']",
//	//		//"price" : crawler.Query(func(p *requestor.Page) string {
//	//		//	o := p.Document().Find("div[class*='pricelabel'] strong").First().Text()
//	//		//	return o
//	//		//}),
//	//	},
//	//}
//	//
//	//otomoto := &crawler.Configuration{
//	//	URL: "http://otomoto.pl/osobowe",
//	//	Next: "[class='next abs'] a",
//	//	Object: "a[class*='img-cover']",
//	//	Template: crawler.Template{
//	//		"name" : "header[class*=om-offer-title] div[class=row] h1",
//	//		"cena" : "div[class=price-cell] span[class=om-price]",
//	//		"ficzery" : crawler.EachText(
//	//			"ul[class*='params-list'] li",
//	//			"small",
//	//			"span",
//	//		),
//	//	},
//	//}
//	//
//
//	//immoscout := &crawler.Configuration{
//	//	URL: "http://www.immobilienscout24.de/Suche/S-2/Wohnung-Kauf?enteredFrom=result_list",
//	//	Next: "a[data-is24-qa='paging_bottom_next']",
//	//	Object: "a[class='result-list-entry__brand-title-container']",
//	//	Template: crawler.Template{
//	//		"name" : "#expose-title",
//	//		"price": "div[class*='is24qa-kaufpreis is24-value']",
//	//		"type" : "dd[class*='is24qa-wohnungstyp']",
//	//		"living_space" : "div[class*='is24qa-wohnflaeche-ca is24-value']",
//	//		"rooms": "div[class*='is24qa-zi is24-value']",
//	//		"description" : "pre[class=is24qa-objektbeschreibung]",
//	//	},
//	//}
//
//	homegate := &Configuration{
//		URL: "http://www.homegate.ch/buy/apartment/matching-list?lastMap=ctn_gr&ab=G000000000000000000000000000000000000000000000000000000000000000000000000000000120000000008E00200C082640001000000000000",
//		Next: "a[rel=next]",
//		Object: "a[class='detail-page-link box-row--link']",
//		//Template: Template{
//		//	"id" : First("div[class=nr] span"),
//		//	"name" : First("h1.title"),
//		//	"price" : First("span[itemprop='price']"),
//		//	"address": First("div[class*=detail-address]"),
//		//	"type" : First("ul[class='list--plain list--flat list--spaced-double text--small'] span[class*='float-right']"),
//		//	"features" : Each(
//		//		"ul[class='list--plain list--flat list--spaced-double text--small'] li",
//		//		"span[class='text--small']",
//		//		"span[class*='float-right']",
//		//	),
//		//	//"description" : First("div[class='description-content']"),
//		//},
//	}
//	//
//	//New().
//	////Add(olx).
//	////Add(immoscout, immoReceiver).
//	////Add(otomoto).
//	//Add(homegate, homeGateReceiver).
//	//Run()
//
//	New().Add(homegate, Homegate{}).Run()
//
//	homeGateTag := tag.New(Homegate{})
//	homeGateTag.AddRule("scrape", ">", ":")
//	homeGateTag.AddRule("filter", ">", ":")
//
//	//method, params := homeGateTag.Get("filter", "Kind", 0)
//	//method, params := f.Get("Kind", 0)
//	//fmt.Println(method, params)
//	//fmt.Println(f)
//	//fmt.Println(t.Get("Test", 1))
//
//
//	//fmt.Println(filter.Alpha("Number of Rooms' 13"))
//
//}
//
//type Homegate struct {
//	Id    int64 `filter:"numeric>" scrape:"first>div[class=nr] span"`
//	Name  string `filter:"text>" scrape:"first>h1.title"`
//	Kind  string `filter:"alpha>" scrape:"first>ul[class='list--plain list--flat list--spaced-double text--small'] span[class*='float-right']"`
//	Price string `filter:"numeric>" scrape:"first>span[itemprop='price']"`
//	//description string
//	//features    HomegateFeatures
//}
//
//type HomegateFeatures struct {
//	rooms        float32
//	availability string
//	livingSpace  int
//}
////
//func (this *Homegate)  UnmarshalJSON(b []byte) (err error) {
//	err = nil
//	fmt.Println(b)
//	return
//}
//
//func homeGateReceiver(record map[string]interface{}) {
//	js, _ := json.Marshal(record)
//
//	o := &Homegate{}
//	err := json.Unmarshal(js, o)
//	if (err != nil) {
//		fmt.Println(err)
//	}
//
//	fmt.Printf("%s \n %s \n\n", js, o)
//}