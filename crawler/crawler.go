package crawler

import (
	"net/url"
	"github.com/sokool/scraper/requestor"
	"fmt"
	"reflect"
	"github.com/kr/pretty"
	"github.com/PuerkitoBio/goquery"
	"github.com/bfontaine/gostruct"
	"github.com/leebenson/conform"
	"os"
	"encoding/xml"
)

type storage struct {
	objects []*Object
}

func (s *storage) put(o Object) *storage {
	s.objects = append(s.objects, &o)

	return s
}

func (s *storage) flush(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "<%s>\n", "objects")
	for _, object := range s.objects {
		bytes, _ := xml.MarshalIndent(object, "", "  ")
		fmt.Fprintf(file, "%s\n", string(bytes))
	}
	fmt.Fprintf(file, "</%s>\n", "objects")
}

type Object interface{}

type Feeder struct {
	Name        string
	Url         string
	Next        string
	Landing     string
	Data        Object
	Destination string
	OnData      func(Object)

	storage     *storage
	dataType    reflect.Type
}

type Crawler struct {
	feeders []*Feeder
	request *requestor.Client
}

func New() *Crawler {
	return &Crawler{
		request: requestor.NewGoQueryClient(),
		feeders: make([]*Feeder, 0),
	}
}

func (c *Crawler) Add(feed *Feeder) *Crawler {
	feed.dataType = reflect.TypeOf(feed.Data)
	feed.storage = &storage{objects:make([]*Object, 0)}
	c.feeders = append(c.feeders, feed)
	return c
}

func (c *Crawler) Scrape() {

	for _, feed := range c.feeders {
		c.page(feed.Url, feed)
	}

	c.request.WaitForAll()

	for _, feed := range c.feeders {
		feed.storage.flush(feed.Name + ".xml")
	}
}

func (f *Feeder) destination(o Object) {
	if f.Destination == "xml" {
		f.storage.put(o)
	}
}

func (c *Crawler) row(item *goquery.Selection, feed *Feeder) {
	uri, _ := item.Attr("href")
	c.request.Do(fqdn(uri, feed.Url), func(result requestor.Result) {

		value := reflect.New(feed.dataType)
		initializeStruct(feed.dataType, value.Elem())
		object := value.Interface()

		gostruct.Populate(object, result.(*goquery.Document))
		conform.Strings(object)

		feed.destination(object)
		feed.OnData(object)

	})
}

func (c *Crawler) page(url string, feed *Feeder) {
	c.request.Do(fqdn(url, feed.Url), func(result requestor.Result) {
		document := result.(*goquery.Document)
		nextUri, ok := document.Find(feed.Next).Attr("href")
		document.Find(feed.Landing).Each(func(i int, item *goquery.Selection) {
			c.row(item, feed)
		})

		if ok {
			c.page(nextUri, feed)
		}
	})
}

func Show(i interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(i))
}

func fqdn(urlA, urlB string) string {
	a, _ := url.Parse(urlA)
	b, _ := url.Parse(urlB)

	if a.Host != "" {
		return urlA
	}
	return fmt.Sprintf("%s://%s%s", b.Scheme, b.Host, a)

}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}