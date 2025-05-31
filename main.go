package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/shafnybuilds/femProject/internal/app"
	"github.com/shafnybuilds/femProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	// in Go unique way of closing the DB connection
	defer app.DB.Close()

	// going to bind the function HealthChcek
	// in go functions are first class citizens, we can pass it like variables
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
