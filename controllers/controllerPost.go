package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gregpmillr/rest-api/models"
	"net/http"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Post: %+v", post)
}
