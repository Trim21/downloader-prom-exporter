package main

import (
	"fmt"
	"net/http"
	"os"

	// This would automatically load/inject environment variables from a .env file
	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/trim21/errgo"

	"app/reporter"
)

func main() {
	if err := Start(); err != nil {
		log.Fatal().Msg("failed to start HTTP server")
	}
}

func Start() error {
	err := reporter.SetupMetrics()
	if err != nil {
		return err
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.Handle("/metrics", promhttp.Handler())

	return errgo.Wrap(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), http.DefaultServeMux), "failed to start http server")
}
