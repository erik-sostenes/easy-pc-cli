package colly

import (
	"github.com/gocolly/colly"
	"log"
)

// NewCollyClient method that returns a colly.Collector
func NewCollyClient() *colly.Collector {
	c := colly.NewCollector(colly.Async(true))

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Add("Keep-Alive", "timeout=100, max=1000")
		log.Println("Visiting, ", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		log.Println("Response Code: ", response.StatusCode)
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Println("error", err.Error())
	})

	return c
}
