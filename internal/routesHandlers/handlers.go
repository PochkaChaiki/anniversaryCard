package routeshandlers

import (
	"html/template"
	"image/png"
	"log"
	"net/http"
	dbmanager "server/internal/dbManager"
	imageprocessor "server/internal/imageProcessor"
)

func GetStaticHandler(address *string, port *string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	cardText, err := dbmanager.GetString()
	if err != nil {
		log.Fatalf("error db getstring: %v\n", err)
		return
	}
	img, err := imageprocessor.CreateImage(cardText)
	if err != nil {
		log.Fatalf("error create image: %v", err)
		return
	}
	png.Encode(w, img)
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
