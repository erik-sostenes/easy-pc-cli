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
	query  application.OfferQuery
	finder ports.OfferFinder[application.OfferQuery]
}

// NewOfferFlags initializes an instance of OfferFlags all its flags assigned
func NewOfferFlags(finder application.OfferFinder) *OfferFlags {
	offer := &OfferFlags{
		FlagSet: flag.NewFlagSet("website", flag.ExitOnError),
		finder:  &finder,
	}

	website := offer.FlagSet

	website.Var(Validate{&offer.query.Website}, "website", "name of the website")
	website.Var(Validate{&offer.query.Container}, "container-offers", "query selector indicating the container of the offer set")
	website.Var(Validate{&offer.query.Item}, "item-offer", "query selector indicating the offer item")
	website.Var(Validate{&offer.query.CategoryId}, "category-id", "id of the website category")
	website.Var(Validate{&offer.query.Category}, "category", "name of the website category")
	website.Var(Validate{&offer.query.Title}, "title", "query selector to search for offer title")
	website.Var(Validate{&offer.query.OriginalPrice}, "original-price", "query selector to search the original price of the offer")
	website.Var(Validate{&offer.query.Discounted}, "discounted", "query selector to search for the offer discounted price")
	website.Var(Validate{&offer.query.Percentage}, "percentage", "query selector to search for offer percentage")
	website.Var(Validate{&offer.query.OfferUrl}, "offer-url", "query selector to search for the offer url 'href'")
	website.Var(Validate{&offer.query.OfferDay}, "offer-day", "query selector to search if it is the offer of the day")
	website.Var(Validate{&offer.query.Available}, "available", "query selector to search for offer availability")
	website.Var(Validate{&offer.query.Delivery}, "delivery", "query selector to search out if the offer is free for delivery ")

	if err := offer.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	return offer
}

// Name method that returns the name of the main command
func (o *OfferFlags) Name() string {
	return o.FlagSet.Name()
}

// Parse method that initializes the whole set of flags of the main command
func (o *OfferFlags) Parse(args []string) error {
	return o.FlagSet.Parse(args)
}

// Run method that executes the task with the values of each flag
func (o *OfferFlags) Run() error {
	o.query.Urls = o.FlagSet.Args()
	return o.finder.Find(o.query)
}
