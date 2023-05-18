package ports

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/domain/bus/query"
)

type (
	// OfferDetailsFinder interface that represents a port that defines the contracts of how the adapters communicate
	// with the core business
	OfferDetailsFinder[V query.Query] interface {
		// Find method that receives a query to validate the data and send the data to the port that will communicate
		// with the find adapter offers
		Find(V) error
	}

	// OfferDetailsScraper interface that represents an output port that defines the contracts of how the adapters communicate
	OfferDetailsScraper interface {
		// Scraping method that receives a data to search the data of an offer by colly
		Scraping(string, string) (domain.OfferDetails, error)
	}

	// HttpRequester interface that represents an output port that defines the contracts of how the adapters communicate
	// port in charge of making requests to the set of injected endpoints at runtime
	HttpRequester interface {
		// Request method that receives a domain.OffersDetails to send the request the set of endpoints
		Request(domain.OfferDetails) error
	}
)
