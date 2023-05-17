package cli

import (
	"flag"
	a "github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/application"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/domain/ports"
	"os"
)

// OfferDetailsFlags represents a set of flags that need the offer details with all required dependencies
type OfferDetailsFlags struct {
	*flag.FlagSet
	a.OfferDetailsQuery
	ports.OfferDetailsFinder[a.OfferDetailsQuery]
}

// NewOfferDetailsFlags initializes an instance of OfferDetailsFlags all its flags assigned
// and injects the required dependencies
func NewOfferDetailsFlags(finder ports.OfferDetailsFinder[a.OfferDetailsQuery]) *OfferDetailsFlags {
	offerDetails := &OfferDetailsFlags{
		FlagSet:            flag.NewFlagSet("offer-details", flag.ExitOnError),
		OfferDetailsFinder: finder,
	}

	command := offerDetails.FlagSet

	command.Var(Validate{&offerDetails.Url}, "url", "inject the url of the offer")
	command.Var(Validate{&offerDetails.Id}, "id", "inject the offer identifier")

	if err := offerDetails.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	return offerDetails
}

// Name method that returns the name of the main command
func (o *OfferDetailsFlags) Name() string {
	return o.FlagSet.Name()
}

// Parse method that initializes the whole set of flags of the main command
func (o *OfferDetailsFlags) Parse(args []string) error {
	return o.FlagSet.Parse(args)
}

// Run method that executes the task with the values of each flag
func (o *OfferDetailsFlags) Run() error {
	return o.OfferDetailsFinder.Find(o.OfferDetailsQuery)
}
