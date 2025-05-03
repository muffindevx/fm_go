package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func New() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}

func (*Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is available\n")
}
