package main

import (
	"os"

	"github.com/brunetto/lost"
	"github.com/sirupsen/logrus"
)

func main() {
	hn, err := os.Hostname()
	if err != nil {
		logrus.Fatal(err)
	}

	l := logrus.New().WithFields(logrus.Fields{
		"app":      "example",
		"hostname": hn,
	})

	lost.LogStartAndStop("example", l)

	for i := 0; ; i++ {
		if i == 5 {
			p, err := os.FindProcess(os.Getpid())
			if err != nil {
				logrus.Fatal(err)
			}
			p.Signal(os.Interrupt)
		}
	}
}
