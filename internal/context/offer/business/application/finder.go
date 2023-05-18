package application

import (
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/domain/bus/query"
	"net/url"
)

// OfferQuery implements the query.Query interface
var _ query.Query = &OfferQuery{}

type OfferQuery struct {
	Website    string
	CategoryId string
	Category   string
	Urls       []string
}

// QueryId method that returns the identifier of OfferQuery
func (o OfferQuery) QueryId() string {
	return "find_offer_query"
}

// OfferFinder implements the ports.OfferFinder interface
var _ ports.OfferFinder[OfferQuery] = &OfferFinder{}

type OfferFinder struct {
	ports.OfferScraper
	ports.HttpRequester
	values map[string]string
}

func NewOfferFinder(scraper ports.OfferScraper, requester ports.HttpRequester) OfferFinder {
	return OfferFinder{
		OfferScraper:  scraper,
		HttpRequester: requester,
		values:        make(map[string]string),
	}
}

// Find method that receives a query, validates the data and communicates with the ports.OfferScraper port
func (o *OfferFinder) Find(query OfferQuery) error {
	if err := o.ensureUrlIsValid(query.Urls); err != nil {
		return err
	}

	o.values["website"] = query.Website
	o.values["category_id"] = query.CategoryId
	o.values["category"] = query.Category

	offers, err := o.OfferScraper.Scraping(o.values, query.Urls)
	if err != nil {
		return err
	}

	return o.HttpRequester.Request(offers)
}

func (o *OfferFinder) ensureUrlIsValid(urls []string) error {
	for _, v := range urls {
		if _, err := url.ParseRequestURI(v); err != nil {
			return fmt.Errorf("offer %s is not valid", v)
		}
	}
	return nil
}
