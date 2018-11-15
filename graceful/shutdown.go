package graceful

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mvignesh1950/listen-me/blogger"
)

/**
 * Handles graceful shutdown.
 */
func ActivateGracefulShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for sig := range signalChan {
		if sig == syscall.SIGINT {
			blogger.Warn().Command("interrupt").Message("Exiting.").Log()
			os.Exit(0)
		}
	}
}
