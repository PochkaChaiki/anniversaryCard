package main

import (
	"fmt"
	dbmanager "server/dbManager"
	img "server/imageProcessor"

	"image/png"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	mux := http.NewServeMux()

	imageHandler := http.HandlerFunc(getImageHandler)

	mux.Handle("/image", enableCORS(imageHandler))
	mux.Handle("/", enableCORS(http.FileServer(http.Dir("./static"))))

	log.Printf("Server listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
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
