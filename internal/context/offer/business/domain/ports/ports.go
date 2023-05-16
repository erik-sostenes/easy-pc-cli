package ports

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/domain/bus/query"
)

type (
	// OfferFinder interface that represents a port that defines the contracts of how the adapters communicate
	// with the core business
	OfferFinder[V query.Query] interface {
		// Find method that receives a query to validate the data and send the data to the port that will communicate
		// with the find adapter offers
		Find(V) error
	}

	// OfferScraper interface that represents an output port that defines the contracts of how the adapters communicate
	OfferScraper interface {
		// Scraping method that receives a domain.offer with query selectors to search the data of an offer by colly
		Scraping(map[string]string, []string) (domain.Offers, error)
	}
)
