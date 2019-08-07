package ps1_fun

import (
	"context"
	"os"
	"os/signal"
)

func ContextWithCancelSignals(sig ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, sig...)
	go func() {
		<-exit
		cancel()
	}()
	return ctx, cancel
}
