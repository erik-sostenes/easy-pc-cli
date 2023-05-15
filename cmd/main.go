package main

import (
	"github.com/erik-sostenes/easy-pc-cli/cmd/cli/bootstrap"
	"log"
	"os"
)

func main() {
	if err := bootstrap.Execute(os.Args[1:]); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
