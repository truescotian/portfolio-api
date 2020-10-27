package controllers

import (
	"encoding/json"
	"github.com/gregpmillr/rest-api/models"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	model := models.NewModel(1, "Greg")
	out, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}
