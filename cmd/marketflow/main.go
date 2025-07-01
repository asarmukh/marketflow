package marketflow

import (
	"context"
	"flag"
	"marketflow/helper"
	"marketflow/internal/util"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// port := flag.Int("port", 8080, "Listening port number")
	help := flag.Bool("help", false, "Show help")

	if *help {
		helper.PrintHelp()
		return
	}

	util.InitLogger()
	util.Logger.Info("Logger initializated")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	util.Logger.Info("Application started")

	<-ctx.Done()
	util.Logger.Info("Shutting down gracefully")
	time.Sleep(1 * time.Second)
	util.Logger.Info("cleanup complete")
}
