package cli

import (
	"flag"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/application"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/domain/ports"
	"os"
)

// OfferFlags represents a set of flags that need the product
type OfferFlags struct {
	*flag.FlagSet
	setFlags struct {
		website       string
		category      string
		title         string
		originalPrice string
		discounted    string
		percentage    string
		offerUrl      string
		offerDay      string
		available     string
		delivery      string
		urls          []string
	}
	finder ports.OfferFinder[application.OfferQuery]
}

// NewProductFlags initializes an instance of ProductsFlags all its flags assigned
func NewProductFlags(finder application.OfferFinder) *OfferFlags {
	offer := &OfferFlags{
		FlagSet: flag.NewFlagSet("website", flag.ExitOnError),
		finder:  &finder,
	}

	website := offer.FlagSet

	website.Var(Validate{&offer.setFlags.website}, "website", "name of the website")
	website.Var(Validate{&offer.setFlags.category}, "category", "name of the website category")
	website.Var(Validate{&offer.setFlags.title}, "title", "query selector to search for product title")
	website.Var(Validate{&offer.setFlags.originalPrice}, "original-price", "query selector to search the original price of the product")
	website.Var(Validate{&offer.setFlags.discounted}, "discounted", "query selector to search for the product discounted price")
	website.Var(Validate{&offer.setFlags.percentage}, "percentage", "query selector to search for product percentage")
	website.Var(Validate{&offer.setFlags.offerUrl}, "offer-url", "query selector to search for the product offer url")
	website.Var(Validate{&offer.setFlags.offerDay}, "offer-day", "query selector to search if it is the product offer of the day")
	website.Var(Validate{&offer.setFlags.available}, "available", "query selector to search for product availability")
	website.Var(Validate{&offer.setFlags.delivery}, "delivery", "query selector to search out if the product is free for delivery ")

	if err := offer.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	return offer
}

// Name method that returns the name of the main command
func (p *OfferFlags) Name() string {
	return p.FlagSet.Name()
}

// Parse method that initializes the whole set of flags of the main command
func (p *OfferFlags) Parse(args []string) error {
	return p.FlagSet.Parse(args)
}

// Run method that executes the task with the values of each flag
func (p *OfferFlags) Run() error {
	query := application.OfferQuery{
		Website:       p.setFlags.website,
		Category:      p.setFlags.category,
		Title:         p.setFlags.title,
		OriginalPrice: p.setFlags.originalPrice,
		Discounted:    p.setFlags.discounted,
		Percentage:    p.setFlags.percentage,
		OfferUrl:      p.setFlags.offerUrl,
		OfferDay:      p.setFlags.offerDay,
		Available:     p.setFlags.available,
		Delivery:      p.setFlags.delivery,
		Urls:          p.FlagSet.Args(),
	}

	return p.finder.Find(query)
}
