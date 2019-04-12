package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	ratelimiter "github.com/mantika/go-ratelimiter"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Register the endpoint
	router.
		Methods("GET").
		Path("/").
		Name("helloWorld").
		HandlerFunc(helloWorld)

	// Register the Ratelimiter
	m := ratelimiter.NewGlobal().
		WithDefaultQuota(4, time.Second).
		Middleware()

	// Attach the Ratelimiter to the router - which in turn contains our endpoint
	n := negroni.New()
	n.Use(m)
	n.UseHandler(router)

	// Start the WebServer
	log.Println("Serving API on port 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}

func helloWorld(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)

	log.Println("Successfully served request")
	rw.WriteHeader(http.StatusOK)
}
