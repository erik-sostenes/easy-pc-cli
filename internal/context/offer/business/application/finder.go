package application

import (
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/domain/bus/query"
	"regexp"
)

// OfferQuery implements the query.Query interface
var _ query.Query = &OfferQuery{}

type OfferQuery struct {
	Website       string
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

type OfferFinder struct{}

func NewOfferFinder() OfferFinder {
	return OfferFinder{}
}

// Find method that receives a query, validates the data and communicates with the ports.OfferScraper port
func (o *OfferFinder) Find(query OfferQuery) (err error) {
	if ok, err := o.ensureUrlIsValid(query.Urls); !ok {
		return err
	}
	return
}

func (o *OfferFinder) ensureUrlIsValid(urls []string) (bool, error) {
	regexExp := "(?:https?:\\/\\/)?(?:[\\w]+\\.)(?:\\.?[\\w]{2,})(\\/[\\w]*)*(\\.[\\w]+)"

	for _, url := range urls {
		if matched, err := regexp.MatchString(regexExp, url); err != nil || !matched {
			return false, fmt.Errorf("offer url %s is not valid", url)
		}
	}

	return true, nil
}
