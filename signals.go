package sig

import (
	"os"
	"os/signal"
	"syscall"
)

type signalCallback func(os.Signal)

func Observe() {
	ObserveWith(func(os.Signal) { os.Exit(0) }, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func ObserveWith(cb signalCallback, sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	go func() {
		cb(<-c)
	}()
}
