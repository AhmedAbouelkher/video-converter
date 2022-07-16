package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AhmedAbouelkher/video-converter/pkg/controllers"
	"github.com/gorilla/mux"
)

// Upload and Save video
// Convert video
// Return playable m3u8 file link

const (
	port = 8005
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HandleDisplayingIndexPage).Methods(http.MethodGet)
	router.HandleFunc("/upload", controllers.HandleUploadingFiled).Methods(http.MethodPost)

	// fs := &serve.JustFilesFilesystem{http.Dir("./storage/compressions")}
	router.PathPrefix("/compressions").Handler(http.StripPrefix("/compressions", http.FileServer(http.Dir("./storage/compressions")))).Methods(http.MethodGet, http.MethodOptions)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
	}
	log.Printf("Starting the server on [http://localhost:%d]\n", port)
	log.Fatal(srv.ListenAndServe())
}
