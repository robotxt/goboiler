package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/robotxt/goboiler/app"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	// Only log the warning severity or above.
	env := strings.ToUpper(os.Getenv("ENVIRONMENT"))

	if env == "PRODUCTION" {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.DebugLevel)
	} else {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// set default port if empty
		port = "9090"
	}

	var router *mux.Router = mux.NewRouter().StrictSlash(false)

	app := app.InitializeApp(context.Background(), router, logger)
	app.ActivateAPI()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second, // timeout in data read
		WriteTimeout: 10 * time.Second, // timeout for response
	}

	go func() {
		logger.Info("Starting Server...")
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()
	waitForShutdown(srv)

}

func checkErr(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	os.Exit(0)
}
