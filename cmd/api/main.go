package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lhervi/blog-api/internal/blog"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")

	r.HandleFunc("/post", blog.CreatePostHandler).Methods("GET")
	r.HandleFun("/post", blog.GetPostHandler).Methods("POST")

	fmt.Println("servidor escuchando por el puerto 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API del Blog")
}
