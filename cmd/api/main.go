package main

import (
	"flag"
	"log/slog"
	"os"
	"strconv"
)

// Declare a string containing the application version number. Later in the book we'll
// generate this automatically at build time, but for now we'll just store the version
// number as a hard-coded global constant.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our application.
// For now, the only configuration settings will be the network port that we want the
// server to listen on, and the name of the current operating environment for the
// application (development, staging, production, etc.). We will read in these
// configuration settings from command-line flags when the application starts.
type config struct {
	port int
	env  string
}

// Define an application struct to hold the dependencies for our HTTP handlers, helpers,
// and middleware. At the moment this only contains a copy of the config struct and a
// logger, but it will grow to include a lot more as our build progresses.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	//an instance of the config struct.
	var cfg config

	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new structured logger which writes log entries to the standard out
	// stream.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the application struct, containing the config struct and
	// the logger.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// This delegates all routing logic (middleware, groups, handlers) to routes.go
	router := app.routes()

	// We must log that the server is starting *before* we actually start it.
	logger.Info("starting server", "addr", cfg.port, "env", cfg.env)

	// Start the server using Gin's built-in run method.
	// This replaces the explicit &http.Server{...} block.
	err := router.Run(":" + strconv.Itoa(cfg.port))

	// If router.Run() returns, it means the server has stopped or failed to start.
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
