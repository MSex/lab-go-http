package main

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
	zlogger, err := injectLogger()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer func() {
		zlogger.Info("Exiting")
		zlogger.Sync()
	}()

	zlogger.Warn("Initializing")

	router, err := inject()
	if err != nil {
		err = errors.Wrap(err, "Error initializing")
		zlogger.DPanic(err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
	//TODO graceful shutdown

}
