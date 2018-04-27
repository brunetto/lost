package lost

import (
	"os"
	"os/signal"
	"syscall"
)

// Logger is a placeholder for the minimum logger interface
type Logger interface {
	Printf(format string, args ...interface{})
}

// LogStartAndStop waits the quit signal
func LogStartAndStop(processName string, logger Logger) {
	// create signal channel
	c := make(chan os.Signal, 1)
	// catch stop signals and send them to the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// spin goroutine reacting to the stop signal
	go func(c chan os.Signal, processName string, logger Logger) {
		// waiting for exit signal on the channel
		<-c
		logger.Printf("%v: stopped by the user", processName)
		os.Exit(0)
	}(c, processName, logger)

	// log process start
	logger.Printf("%v: started", processName)
}
