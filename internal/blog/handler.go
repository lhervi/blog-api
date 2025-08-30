package blog

import (
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post creado con éxito"))
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Lista de Post"))
}
