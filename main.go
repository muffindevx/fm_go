package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/muffindevx/fm_go/internal/app"
	"github.com/muffindevx/fm_go/internal/routes"
)

func main() {
	application, err := app.New()
	// Allow to get the variable from console, in this case, we pass the port for our service
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	if err != nil {
		// We use panic if we do not expect an error
		panic(err)
	}

	r := routes.Setup(application)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
		/*
			The maximum amount of time to wait for the next request when keep-alive is enabled
			If IdleTimeout is zero, the value of ReadTimeout is used
		*/
		IdleTimeout: time.Minute,
		/*
			Maximum duration for reading the entire request, including the body
		*/
		ReadTimeout: 10 * time.Second,
		/*
			Maximum duration before timing out writes of the response
		*/
		WriteTimeout: 30 * time.Second,
	}

	application.Logger.Printf("We are running on port %d\n", port)

	if err := server.ListenAndServe(); err != nil {
		application.Logger.Fatal(err)
	}

}
