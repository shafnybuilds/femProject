package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/shafnybuilds/femProject/internal/api"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// our stores will go here

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	// application instance
	app := &Application{
		Logger: logger,
		WorkoutHandler: workoutHandler,
	}

	// nil are valid error types, they satisfy the interface for error
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available \n")
}
