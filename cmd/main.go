package main

import (
	"fmt"
	dbmanager "server/internal/dbManager"
	img "server/internal/imageProcessor"

	"flag"
	"html/template"
	"image/png"
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

	imageHandler := http.HandlerFunc(getImageHandler)
	staticHandler := http.HandlerFunc(getStaticHandler)

	mux.Handle("/image", enableCORS(imageHandler))
	mux.Handle("/", enableCORS(staticHandler))

	log.Printf("Server listening on %s:%s", *address, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *address, *port), mux))
}

func getStaticHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.New("index.html").ParseFiles("./static/index.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	data := struct {
		IPAddress string
		Port      string
	}{
		IPAddress: *address,
		Port:      *port,
	}
	err = templ.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	cardText, err := dbmanager.GetString()
	if err != nil {
		log.Fatalf("error db getstring: %v\n", err)
		return
	}
	img, err := img.CreateImage(cardText)
	if err != nil {
		log.Fatalf("error create image: %v", err)
		return
	}
	png.Encode(w, img)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
