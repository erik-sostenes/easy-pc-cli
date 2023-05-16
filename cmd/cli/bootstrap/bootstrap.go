// bootstrap package that builds the program with its full set of components

package bootstrap

import (
	"errors"
	"fmt"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/application"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/input/cli"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/output/scraping"
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

	colly := colly.NewCollyClient()
	offerFinder := application.NewOfferFinder(scraping.NewOfferScraper(*colly))

	switch os.Args[1] {
	case "website":
		offer := cli.NewOfferFlags(offerFinder)
		if err := offer.Run(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%s command no fount", os.Args[1])
	}
	return nil
}
