package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var posts = []Post{}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var newPost Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Generar un ID simple para el nuevo post
	newPost.ID = fmt.Sprint("post-%d", len(posts)+1)
	newPost.CreatedAt = time.Now()

	//Añadir el nuevo post a la lista
	posts = append(posts, newPost)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request)
{
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Lista de Post"))
}
