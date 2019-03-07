package main

import (
	"flag"
	"net"
	"net/http"

	"github.com/prestonp/goservice/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	debug = flag.Bool("debug", false, "sets log level to debug")
	host  = flag.String("host", "", "host")
	port  = flag.String("port", "3000", "port")
)

func main() {
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("log level set to debug")

	svc := service.New()
	addr := net.JoinHostPort(*host, *port)
	panic(http.ListenAndServe(addr, svc))
}
