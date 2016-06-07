package requestor

import (
	query "github.com/PuerkitoBio/goquery"
	"sync"
	"net/url"
	"log"
	"os"
	"github.com/bfontaine/gostruct"
	"fmt"
)

type Client struct {
	onRequest caller
	onFinish  finisher
}
type request struct {
	url      *url.URL
	onResult Response
}

type Response func(Result)

type Engine func(*url.URL) Result

type caller func(*request)

type finisher func()

type Result interface{}

func (c *Client) Do(httpUrl string, onResponse Response) {
	link, ok := url.Parse(httpUrl)
	if (ok != nil) {
		panic(ok)
	}

	c.onRequest(&request{link, onResponse})
}

func (c *Client) WaitForAll() {
	c.onFinish()
}

// Create framework in order to make async calls. It returns two functions:
// Caller: function takes Request and put it on the queue
// Finisher: function to let framework know that no more Request are going to be created.
func server(engine Engine) (caller, finisher) {
	delay := sync.WaitGroup{}
	requests := make(chan *request)
	go func(invokers <- chan *request) {
		for invoker := range invokers {
			delay.Add(1)
			go func(request *request) {
				request.onResult(engine(request.url))
				defer delay.Done()
			}(invoker)
		}
	}(requests)

	return caller(func(r *request) {
		requests <- r
	}), finisher(func() {
		delay.Wait()
	})
}

func NewClient(e Engine) *Client {
	caller, finisher := server(e)
	return &Client{
		onRequest: caller,
		onFinish: finisher,
	}
}

func NewGoQueryClient() *Client {
	return NewClient(func(u *url.URL) Result {
		d, err := query.NewDocument(u.String())
		if err != nil {
			log.Fatalf("\t %s %s\n", err, u)
		}

		return d
	})
}

func NewGoStruct(target interface{}) *Client {

	return NewClient(func(u *url.URL) Result {
		err := gostruct.Fetch(target, u.String())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		return target
	})

}

//
//sites := []string{
//"http://google.pl",
//"http://onet.pl",
//"http://wp.pl",
//"http://polsat.pl",
//"http://tvn.pl",
//"http://facebook.pl",
//"http://fakty.interia.pl/polska",
//"http://mint-soft.pl",
//"http://google.de",
//"http://www.wroclaw.apodatkowa.gov.pl/izba-skarbowa-we-wroclawiu;jsessionid=C3FE209C723806A6F34547C5EF3C5C34",
//"http://stackoverflow.com",
//"http://allegro.pl",
//"http://olx.pl",
//"http://otomoto.pl",
//"http://otodom.pl",
//"http://gradka.pl",
//"http://eurosport.onet.pl",
//}
//
//documentResponse := func(v requestor.Result) {
//	document := v.(*goquery.Document)
//	a, _ := document.Find("meta[name=description]").Attr("content")
//	if a == "" {
//		a = document.Find("title").Text()
//	}
//
//	fmt.Printf("%s --> %# v\n", document.Url.String(), a)
//}
//
//r := requestor.NewGoQuery()
//for i, url := range sites {
//r.Do(url, documentResponse)
//if i == 0 {
//time.Sleep(time.Second)
//}
//}
//r.WaitForAll()