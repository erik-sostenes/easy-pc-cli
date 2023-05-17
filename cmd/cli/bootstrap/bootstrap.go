// bootstrap package that builds the program with its full set of components

package bootstrap

import (
	"errors"
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/application"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/input/cli"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/output/http"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/output/scraping"
	a2 "github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/business/application"
	c2 "github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/infrastructure/input/cli"
	http2 "github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/infrastructure/output/http"
	s2 "github.com/erik-sostenes/easy-pc-cli/internal/context/offer_details/infrastructure/output/scraping"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/shared/infrastruture/colly"
	"os"
)

// Runner will define all subcommands and will initialize
type Runner interface {
	// Parse parses flag definitions from the argument list,
	// does not include the command name, only the subcommands
	Parse([]string) error
	// Run execute the process with the values of the subcommands
	Run() error
	// Name returns the name of the subcommand to be worked with
	Name() string
}

// Execute method that initializes the program startup with all dependencies initialized
func Execute(args []string) error {
	if len(args) < 1 {
		return errors.New("website subcommands were expected")
	}

	colly := *colly.NewCollyClient()

	switch os.Args[1] {
	case "offers":
		offerFinder := application.NewOfferFinder(scraping.NewOfferScraper(colly), http.Requester{})
		return cli.NewOfferFlags(offerFinder).Run()
	case "offer-details":
		offerDetails := a2.NewOfferDetailsFinder(s2.NewOfferDetailsScraper(colly), http2.Requester{})
		return c2.NewOfferDetailsFlags(offerDetails).Run()
	default:
		return fmt.Errorf("%s command no fount", os.Args[1])
	}
}
