package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Kialakun/goproxy"
	log "github.com/sirupsen/logrus"
)

var PORT = os.Getenv("PORT")

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	rate, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		description := "ERROR"
		log.WithFields(log.Fields{"Message": "no rate limit specified, setting rate limit to 10MB/s"}).Info(description)
		rate = 10000 * 1024 // 10 MB/s
	}
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go goproxy.Listen(PORT, rate, shutdown)
	<-shutdown
}
