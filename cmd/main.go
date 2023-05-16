package main

import (
	"context"
	"github.com/erik-sostenes/easy-pc-cli/cmd/cli/bootstrap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	exit := make(chan os.Signal, 1)
	doneCH := make(chan struct{}, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-exit
		log.Println("interrupt signal received, waiting for an additional 10 seconds...")
		time.Sleep(time.Second * 5)
		cancel()
	}()

	go func() {
		if err := bootstrap.Execute(os.Args[1:]); err != nil {
			log.Println(err)
		}
		doneCH <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Println("Request is being made to complete the process")
	case <-doneCH:
		log.Println("the process has been completed")
	}

	close(exit)
	close(doneCH)
}
