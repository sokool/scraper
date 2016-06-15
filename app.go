package main

import (
	"github.com/sokool/scraper/thief"
	"runtime"
)

type Property struct {
	Id          string `gostruct:"div.nr span" conform:"trim"`
	Url         string `gostruct:"link[rel=canonical]/href"`
	Name        string `gostruct:"h1.title" conform:"lower,trim,title"`
	Kind        string `gostruct:"div.detail-key-data li:nth-child(1) span:last-child" conform:"trim"`
	Price       string `gostruct:"span[itemprop=price]" conform:"num"`
	Currency    string `gostruct:"span[itemprop=priceCurrency]"`
	Description string `gostruct:"div.detail-description" conform:"trim"`
	Features    Features `gostruct:"div.detail-key-data"`
	Images      string `gostruct:"figure.slick-active img/src"`
}

type Features struct {
	Rooms          string `gostruct:"li:nth-child(2) span:last-child" xml:"Rooms,omitempty"`
	Floor          string `gostruct:"li:nth-child(4) span:last-child" xml:"Floor,omitempty"`
	NumOfFloor     string `gostruct:"li:nth-child(5) span:last-child" xml:"NummberOfFloors,omitempty"`
	LivingSpace    string `gostruct:"li:nth-child(6) span:last-child" xml:"LivingSpace,omitempty"`
	FloorSpace     string `gostruct:"li:nth-child(7) span:last-child" xml:"FloorSpace,omitempty"`
	LotSize        string `gostruct:"li:nth-child(9) span:last-child" xml:"LotSize,omitempty"`
	Volume         string `gostruct:"li:nth-child(10) span:last-child" xml:"Volume,omitempty"`
	RoomHeight     string `gostruct:"li:nth-child(11) span:last-child" xml:"RoomHeight,omitempty"`
	YearOfBuilding string `gostruct:"li:nth-child(17) span:last-child" xml:"YearOfBuild,omitempty"`
	Available      string `gostruct:"li:nth-child(19) span:last-child" xml:"Available,omitempty"`
	LastRenovation string `gostruct:"li:nth-child(18) span:last-child" xml:"LastRenovation,omitempty"`
}

func main() {
	runtime.GOMAXPROCS(4)

	//link := "http://www.homegate.ch/buy/real-estate/canton-graubuenden"
	link := "http://www.homegate.ch/buy/real-estate/region-surselva/matching-list?lastMap=ctn_gr&ep=10"
	homegate := thief.NewTemplate("xml", link, "documents", "next").
	Add("buy", "li.page-nav-item--buy a", "sections").
	Add("sections", "div.ad-aside-content div.box a", "countries").
	Add("countries", "area[href]", "cantons").
	Add("cantons", "area[href]", "documents", "next").
	Add("next", "a[rel=next]", "documents", "next").
	Add("documents", "a[class='detail-page-link box-row--link']", "id", "name", "kind", "price", "currency").
	Add("id", "div.nr span").
	Add("name", "h1.title").
	Add("kind", "div.detail-key-data li:nth-child(1) span:last-child").
	Add("price", "span[itemprop=price]").
	Add("currency", "span[itemprop=priceCurrency]")

	//z500pl := thief.NewTemplate("json", Property{}, "http://z500.pl/domy.html?view=small", "document", "list").
	//Add("list", func(url string) {
	//
	//}).
	//Add("list2", "meta[property=og:url]|replace()").
	//Add("document", "div.projects-listing .short-summary a")

	//ekstradom := thief.NewTemplate("json", Property{}, "http://www.extradom.pl/", "document", "list").
	//Add("list", "ul.pagination a[aria-label=Następna]", "list", "document").
	//Add("document", "div.projects-container div.design-thumb-info a")

	thief.
	New().
	//Add(ekstradom).
	Add(homegate).
	Run()

	//Start("http://www.homegate.ch/buy/real-estate/canton-graubuenden", []string{"cantons"})
	//Start("http://www.homegate.ch/buy/real-estate/switzerland", []string{"countries"})
	//Start("http://www.homegate.ch/kaufen/wohnung/schweiz", []string{"countries"})
	//Start("http://www.homegate.ch/buy/real-estate/region-surselva/matching-list?lastMap=ctn_gr&ep=1", []string{"next", "documents"})
	//Start("http://www.homegate.ch", []string{"buy"})
}
