// bootstrap package that builds the program with its full set of components

package bootstrap

import (
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/business/application"
	"github.com/erik-sostenes/easy-pc-cli/internal/context/offer/infrastructure/input/cli"
	"log"
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
		log.Println("website subcommands were expected")
		os.Exit(1)
	}

	offerFinder := application.NewOfferFinder()

	cmds := []Runner{
		cli.NewProductFlags(offerFinder),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}

	return nil
}
