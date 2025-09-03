package blog

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//var posts = []Post{}

var posts = make(map[string]Post)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var newPost Post

	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Generar un ID simple para el nuevo post
	newPost.ID = fmt.Sprintf("post-%d", len(posts)+1)
	newPost.CreatedAt = time.Now()

	//Añadir el nuevo post a la lista
	//posts = append(posts, newPost)

	posts[newPost.ID] = newPost

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postsID := vars["id"]

	post, ok := posts[postsID]
	if !ok {
		http.Error(w, "Post no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postID := vars["id"]

	var updatedPost Post
	err := json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, ok := posts[postID]
	if !ok {
		http.Error(w, "Post no encontrado", http.StatusNotFound)
		return
	}

	// Mantener el ID original y la fecha de creación
	updatedPost.ID = postID
	updatedPost.CreatedAt = posts[postID].CreatedAt
	posts[postID] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	_, ok := posts[postID]
	if !ok {
		http.Error(w, "Post no encontrado", http.StatusNotFound)
		return
	}

	delete(posts, postID)
	w.WriteHeader(http.StatusNoContent)
}
