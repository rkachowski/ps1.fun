package main

import (
	ps1_fun "github.com/rkachowski/ps1.fun/internal/pkg/ps1-fun"
	"os"
)

func main() {
	ctx, cancel := ps1_fun.ContextWithCancelSignals(os.Interrupt, os.Kill)
	defer cancel()

	errors := make(chan error, 1)
	go func() {
		errors <- ps1_fun.Serve(ctx)
	}()


	select {
	case err := <-errors:
		panic(err)
	case <-ctx.Done():
		//i am cleaning myself up before i die like a nice app
		return
	}
}
