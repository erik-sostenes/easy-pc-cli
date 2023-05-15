package cli

import (
	"errors"
	"flag"
	"strings"
)

// validate that the Value interface of the flags package implements the Validate structure
var _ flag.Value = &Validate{}

// Validate struct that stores a value of a flag to define our own validations
type Validate struct {
	value string
}

// String method that returns the value of flag
func (v Validate) String() string {
	return v.value
}

// Set method that validates the input value of a flag
func (v Validate) Set(value string) (err error) {
	if strings.TrimSpace(value) == "" {
		err = errors.New("unknown subcommand")
		return
	}

	v.value = value
	return
}
