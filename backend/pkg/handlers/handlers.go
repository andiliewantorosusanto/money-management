package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andiliewantorosusanto/money-management/pkg/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the home page")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Number of bytes written: %d", n)
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is the about page")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Number of bytes written: %d", n)
}

func AddValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
	}

	fmt.Fprintf(w, "%f divide by %f is %f", 100.0, 10.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y

	return result, nil
}
