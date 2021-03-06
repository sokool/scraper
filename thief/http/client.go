package http

import (
	query "github.com/PuerkitoBio/goquery"
	"net/http"
	"runtime"
	"time"
	"fmt"
)

var errors, sucess int = 0, 0
var mem uint64 = 0
var start time.Time = time.Now()

func Get(url string) *query.Document {
	res, e := http.Get(url)
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	if e != nil || res.StatusCode != 200 {
		errors++
		return nil
	}
	sucess++

	mem = mem + memStats.Alloc
	since := time.Since(start).Seconds()
	fmt.Printf("\r running:%.2fs, errors: %d, success: %d[avg: %.2f], gorutins: %d, memory: %.2f [avg:%.2f]",
		since,
		errors,
		sucess,
		float64(sucess)/since,
		runtime.NumGoroutine(),
		toMB(memStats.Alloc),
		toMB(uint64(mem/uint64(sucess))),
	)
	document, err := query.NewDocumentFromResponse(res)

	if err != nil {

		return nil
	}

	return document
}

func toMB(mem uint64) float64 {
	return float64(mem) / 1024 / 1024
}
