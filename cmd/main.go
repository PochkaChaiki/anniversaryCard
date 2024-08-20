package main

import (
	"fmt"
	routeshandlers "server/internal/routesHandlers"

	"flag"
	"log"
	"net/http"
)

var (
	address = flag.String("a", "localhost", "IP address of the host")
	port    = flag.String("p", "8000", "Port number of the listening service")
)

// ./server -a <Public IP> -p <Open public port>

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	imageHandler := http.HandlerFunc(routeshandlers.GetImageHandler)
	staticHandler := http.HandlerFunc(routeshandlers.GetStaticHandler(address, port))

	mux.Handle("/image", routeshandlers.EnableCORS(imageHandler))
	mux.Handle("/", routeshandlers.EnableCORS(staticHandler))

	log.Printf("Server listening on %s:%s", *address, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s" /**address,*/, *port), mux))
}
