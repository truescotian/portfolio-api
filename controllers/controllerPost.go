package controllers

import (
	"github.com/truescotian/portfolio-api/ent"

	"encoding/json"
	"fmt"
	"net/http"
)

// CreatePost creates a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	client := ent.Client
	post, err := client.Post.
		Create().
		SetName("truescotian").
		SetFirstName("Greg").
		SetLastName("Miller").
		Save(ent.Context)

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errs := post.Validate(); len(errs) > 0 {
		http.Error(w, fmt.Printf("%v", errs), http.StatusBadRequest)
		return
	}

	if err = post.Create(); err != nil {
		http.Error(w, fmt.Printf("%v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
