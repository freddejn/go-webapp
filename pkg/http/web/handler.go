package web

import (
	"encoding/json"
	"net/http"

	"github.com/freddejn/go-webapp/pkg/domain/listing"
	"github.com/go-martini/martini"
)

func Handler(l listing.Service) http.Handler {
	router := martini.Classic()
	router.Get("/articles", getHome(l))
	return router
}

func getHome(s viewing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		list := s.GetHome()
		json.NewEncoder(w).Encode(list)
	}
}
