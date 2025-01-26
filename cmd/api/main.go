package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// var version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Get config var for terminal
	flag.IntVar(&cfg.port, "port", 4000, "port for the server")
	flag.StringVar(&cfg.env, "enviroment", "development", "development|staging|production")

	flag.Parse()

	// Init logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Init applicaiton instance
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Init server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
