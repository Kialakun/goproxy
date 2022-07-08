package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kialakun/goproxy"
	log "github.com/sirupsen/logrus"
)

var PORT = os.Getenv("PORT")

func CMD() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	var port string
	flag.StringVar(&port, "port", "8888", "Proxy listen port")
	var rate int
	flag.IntVar(&rate, "rate", 512*1024, "Proxy bandwidth ratelimit")

	flag.Parse()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go goproxy.Listen(PORT, rate, shutdown)
	<-shutdown
}
