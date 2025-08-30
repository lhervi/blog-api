package blog

import (
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.writeHeader(http.Statuscreated)
	w.write([]byte("Post creado con éxito"))
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	w.writeHeader(http.StatusOk)
	w.write([]byte("Lista de Post"))
}
