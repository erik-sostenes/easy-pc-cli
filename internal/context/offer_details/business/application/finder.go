package application

import (
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/domain/ports"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/domain/bus/query"
	"net/url"
)

// OfferDetailsQuery implements the query.Query interface
var _ query.Query = OfferDetailsQuery{}

type OfferDetailsQuery struct {
	Url string
	Id  string
}

// QueryId method that returns the identifier of OfferDetailsQuery
func (OfferDetailsQuery) QueryId() string {
	return "find_offer_details_query"
}

// OfferDetailsFinder implements the ports.OfferDetailsFinder interface
var _ ports.OfferDetailsFinder[OfferDetailsQuery] = &OfferDetailsFinder{}

type OfferDetailsFinder struct {
	ports.OfferDetailsScraper
	ports.HttpRequester
}

// NewOfferDetailsFinder returns an instance of OfferDetailsFinder injecting all dependencies
// required by the OfferDetailsFinder container
func NewOfferDetailsFinder(scraper ports.OfferDetailsScraper, requester ports.HttpRequester) *OfferDetailsFinder {
	return &OfferDetailsFinder{
		OfferDetailsScraper: scraper,
		HttpRequester:       requester,
	}
}

// Find method that receives a query, validates the data and communicates with the ports.OfferDetailsScraper port
func (o *OfferDetailsFinder) Find(query OfferDetailsQuery) error {
	if err := o.ensureUrlIsValid(query.Url); err != nil {
		return err
	}

	offerDetails, _ := o.OfferDetailsScraper.Scraping(query.Url, query.Id)

	return o.HttpRequester.Request(offerDetails)
}

// ensureUrlIsValid method that validates the format of the url
func (o *OfferDetailsFinder) ensureUrlIsValid(urls ...string) error {
	for _, v := range urls {
		if _, err := url.ParseRequestURI(v); err != nil {
			return fmt.Errorf("offer v %s is not valid", v)
		}
	}
	return nil
}
