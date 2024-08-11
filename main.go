package main

import (
	"fmt"
	img "server/imageProcessor"

	// "image"
	"image/png"
	"log"
	"net/http"
	// "os"
)

func main() {
	fmt.Println("Starting server")

	// f, err := os.Create("./scrap/txt.png")
	// if err != nil {
	// 	fmt.Printf("Error oqqured: %v", err)
	// 	return
	// }
	// png.Encode(f, img)
	// fmt.Println("Photo made")
	// fmt.Printf("Photo name: %s\n", img)
	http.HandleFunc("/", getImageHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getImageHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	img, err := img.CreateImage("hello, World!")
	if err != nil {
		fmt.Printf("Error oqqured: %v", err)
		return
	}
	png.Encode(w, img)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
