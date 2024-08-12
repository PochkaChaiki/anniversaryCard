package main

import (
	"fmt"
	dbmanager "server/dbManager"
	img "server/imageProcessor"

	// "image"
	"image/png"
	"log"
	"net/http"
	// "os"
)

func main() {
	fmt.Println("Starting server")

	http.HandleFunc("/", getImageHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	cardText, err := dbmanager.GetString()
	if err != nil {
		fmt.Printf("Error oqqured: %v", err)
		return
	}
	img, err := img.CreateImage(cardText)
	if err != nil {
		fmt.Printf("Error oqqured: %v", err)
		return
	}
	png.Encode(w, img)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
