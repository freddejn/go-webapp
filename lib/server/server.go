package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/freddejn/go-webapp/lib/config"
	"github.com/freddejn/go-webapp/lib/handler"
)

func addHandlers(serverMux *http.ServeMux) {
	ah := handler.NewArticleHandlers()
	// serverMux.HandleFunc("/", indexHandler)
	serverMux.HandleFunc("/articles", ah.Handler)
}

func handleRequests() {
	conf := config.GetConfig()
	fmt.Printf(conf.Server.PORT)
	mux := http.NewServeMux()
	addHandlers(mux)
	fmt.Printf(conf.Resources.Templates)
	log.Fatal(http.ListenAndServe(":"+conf.Server.PORT, mux))
}

// StartServer ... Starts the server
func StartServer() {
	handleRequests()
}
