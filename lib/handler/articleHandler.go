package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Article ... Struct for articles
type Article struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

// Articles ... test
var Articles []Article

type articleHandlers struct {
	store map[string]Article
}

func (a *articleHandlers) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		a.get(w, r)
	case "POST":
		a.get(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}

}

func (a *articleHandlers) get(w http.ResponseWriter, r *http.Request) {
	articles := make([]Article, len(a.store))
	i := 0
	for _, article := range a.store {
		articles[i] = article
		i++
	}

	jsonBytes, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (a *articleHandlers) post(w http.ResponseWriter, r *http.Request) {
	// articles := make([]Article, len(a.store))
	fmt.Printf("test")
}

func NewArticleHandlers() *articleHandlers {
	return &articleHandlers{
		store: map[string]Article{
			"id": Article{
				Title:   "Title 1",
				Content: "Content1",
			},
		},
	}
}
