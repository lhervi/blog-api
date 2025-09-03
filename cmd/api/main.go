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
	r.HandleFunc("/posts", blog.GetAllPosts).Methods("GET") // <-- Esta es la que se está llamando
	r.HandleFunc("/posts/{id}", blog.GetPostsHandler).Methods("GET")

	r.HandleFunc("/posts", blog.CreatePostHandler).Methods("POST") // <-- Esta es la que se debería llamar

	r.HandleFunc("/posts/{id}", blog.UpdatePostHandler).Methods("PUT")
	r.HandleFunc("/posts/{id}", blog.DeletePostHandler).Methods("DELETE")

	fmt.Println("servidor escuchando por el puerto 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API del Blog")
}
