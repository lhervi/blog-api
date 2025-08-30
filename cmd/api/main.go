package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"blog-api/internal/blog"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")

	//r.HandleFunc("/posts", blog.CreatePostHandler).Methods("GET")
	//r.HandleFunc("/posts", blog.GetPostsHandler).Methods("POST")

	r.HandleFunc("/posts", blog.GetPostsHandler).Methods("GET")    // <-- Esta es la que se está llamando
	r.HandleFunc("/posts", blog.CreatePostHandler).Methods("POST") // <-- Esta es la que se debería llamar

	fmt.Println("servidor escuchando por el puerto 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API del Blog")
}
