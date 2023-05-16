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
	Website       string
	Container     string
	Item          string
	CategoryId    string
	Category      string
	Title         string
	OriginalPrice string
	Discounted    string
	Percentage    string
	OfferUrl      string
	OfferDay      string
	Available     string
	Delivery      string
	Urls          []string
}

// QueryId method that returns the identifier of OfferQuery
func (o OfferQuery) QueryId() string {
	return "find_offer_query"
}

// OfferFinder implements the ports.OfferFinder interface
var _ ports.OfferFinder[OfferQuery] = &OfferFinder{}

type OfferFinder struct {
	ports.OfferScraper
	queries map[string]string
}

func NewOfferFinder(scraper ports.OfferScraper) OfferFinder {
	return OfferFinder{
		OfferScraper: scraper,
		queries:      make(map[string]string),
	}
}

// Find method that receives a query, validates the data and communicates with the ports.OfferScraper port
func (o *OfferFinder) Find(query OfferQuery) error {
	if err := o.ensureUrlIsValid(query.Urls); err != nil {
		return err
	}

	o.queries["website"] = query.Website
	o.queries["container"] = query.Container
	o.queries["item"] = query.Item
	o.queries["category_id"] = query.CategoryId
	o.queries["category"] = query.Category
	o.queries["title"] = query.Title
	o.queries["original_price"] = query.OriginalPrice
	o.queries["discounted"] = query.Discounted
	o.queries["percentage"] = query.Percentage
	o.queries["offer_url"] = query.OfferUrl
	o.queries["offer_day"] = query.OfferDay
	o.queries["available"] = query.Available
	o.queries["delivery"] = query.Delivery

	_, err := o.OfferScraper.Scraping(o.queries, query.Urls)
	if err != nil {
		return err
	}

	return nil
}

func (o *OfferFinder) ensureUrlIsValid(urls []string) error {
	for _, v := range urls {
		if _, err := url.ParseRequestURI(v); err != nil {
			return fmt.Errorf("offer v %s is not valid", v)
		}
	}

	return nil
}
