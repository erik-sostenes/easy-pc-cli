package cli

import (
	"flag"
	"os"
)

// ProductFlags represents a set of flags that need the product
type ProductFlags struct {
	*flag.FlagSet
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
}

// NewProductFlags initializes an instance of ProductsFlags all its flags assigned
func NewProductFlags() *ProductFlags {
	product := &ProductFlags{
		FlagSet: flag.NewFlagSet("website", flag.ExitOnError),
	}

	website := product.FlagSet

	website.Var(Validate{product.website}, "website-name", "name of the website")
	website.Var(Validate{product.category}, "category", "name of the website category")
	website.Var(Validate{product.title}, "title", "query selector to search for product title")
	website.Var(Validate{product.originalPrice}, "original-price", "query selector to search the original price of the product")
	website.Var(Validate{product.discounted}, "discounted", "query selector to search for the product discounted price")
	website.Var(Validate{product.percentage}, "percentage", "query selector to search for product percentage")
	website.Var(Validate{product.offerUrl}, "offer-url", "query selector to search for the product offer url")
	website.Var(Validate{product.offerDay}, "offer-day", "query selector to search if it is the product offer of the day")
	website.Var(Validate{product.available}, "available", "query selector to search for product availability")
	website.Var(Validate{product.delivery}, "delivery", "query selector to search out if the product is free for delivery ")

	if err := product.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	return product
}

// Name method that returns the name of the main command
func (p *ProductFlags) Name() string {
	return p.FlagSet.Name()
}

// Parse method that initializes the whole set of flags of the main command
func (p *ProductFlags) Parse(args []string) error {
	return p.FlagSet.Parse(args)
}

// Run method that executes the task with the values of each flag
func (p *ProductFlags) Run() error {
	return nil
}
