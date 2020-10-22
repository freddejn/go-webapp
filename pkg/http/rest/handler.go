package rest

import (
	"encoding/json"
	"net/http"

	"github.com/freddejn/go-webapp/pkg/domain/listing"
	"github.com/go-martini/martini"
)

func Handler(l listing.Service) http.Handler {
	router := martini.Classic()
	router.Get("/api/articles", getArticles(l))
	static := martini.Static("pkg/storage/html")
	router.NotFound(static, notFound())
	return router
}

func getArticles(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetAllArticles()
		json.NewEncoder(w).Encode(list)
	}
}

func notFound() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("<html><body>notfound</body></html>"))
	}
}
