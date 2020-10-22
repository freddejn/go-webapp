package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freddejn/go-webapp/pkg/config"
	"github.com/freddejn/go-webapp/pkg/http/rest"
	"github.com/freddejn/go-webapp/pkg/listing"
	"github.com/freddejn/go-webapp/pkg/storage/json"
)

const (
	templatesPath = "templates/"
)

func main() {
	var config = config.GetConfig("dev")
	var lister listing.Service
	s, _ := json.NewStorage()
	lister = listing.NewService(s)
	router := rest.Handler(lister)
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(config.Server.Host+":"+config.Server.PORT, router))
}
