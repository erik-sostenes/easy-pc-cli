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
		FlagSet: flag.NewFlagSet("offers", flag.ExitOnError),
		finder:  &finder,
	}

	command := offer.FlagSet

	command.Var(Validate{&offer.query.Website}, "website", "inject the name of the website")
	command.Var(Validate{&offer.query.CategoryId}, "category-id", "inject the category identifier")
	command.Var(Validate{&offer.query.Category}, "category", "inject the name of the category")

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
