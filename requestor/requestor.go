package requestor

import (
	query "github.com/PuerkitoBio/goquery"
	"sync"
	"net/url"
	"time"
	"log"
)

type invoker struct {
	url    *url.URL
	onLoad func(*Page)
}

type Page struct {
	document    *query.Document // Representation of jQuery framework
	number      int             // Page number
	loadingTime time.Duration   // Time of page fetching
}

type Request struct {
	call func(u *url.URL, f func(*Page))
	done func()
}

func (p *Page) Document() *query.Document {
	return p.document
}

func (p *Page) LoadingTime() time.Duration {
	return p.loadingTime
}

func (r *Request) Do(httpUrl string, response func(*Page)) {
	u, ok := url.Parse(httpUrl)
	if (ok != nil) {
		panic(ok)
	}

	r.call(u, response)
}

func (r *Request) WaitForAll() {
	r.done()
}

func server() (func(*url.URL, func(*Page)), func()) {
	delay := sync.WaitGroup{}
	stream := make(chan *invoker)

	request := func(u *url.URL, response func(*Page)) {
		stream <- &invoker{url: u, onLoad: response}
	}

	finished := func() {
		delay.Wait()
	}

	go func(invokers <- chan *invoker) {
		for i := range invokers {
			delay.Add(1)
			go func(invoker *invoker) {
				defer delay.Done()
				start := time.Now()
				invoker.onLoad(&Page{
					document: fetch(invoker.url),
					loadingTime: time.Since(start),
				})

			}(i)
		}
	}(stream)

	return request, finished
}

func fetch(url *url.URL) *query.Document {
	d, err := query.NewDocument(url.String())
	if err != nil {
		log.Fatalf("\t %s %s\n", err, url)
	}
	//log.Printf("%s\n", url)

	return d

}

func New() *Request {
	start, stop := server()
	return &Request{
		call: start,
		done: stop,
	}
}