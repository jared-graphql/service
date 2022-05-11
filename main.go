package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jaredhughes1012/parcel"
)

type note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	notes := []note{
		{
			Title:       "Title 1",
			Description: "Description 1",
			Date:        time.Now(),
		},
		{
			Title:       "Title 2",
			Description: "Description 2",
			Date:        time.Now(),
		},
		{
			Title:       "Title 3",
			Description: "Description 3",
			Date:        time.Now(),
		},
	}

	_ = parcel.RenderResponse(w, http.StatusOK, &notes)
}

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"*"},
	}))

	r.Get("/notes", handler)

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
