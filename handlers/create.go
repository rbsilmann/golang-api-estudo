package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rbsilmann/api-estudo/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao realizar o decode do JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(todo)
	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   "true",
			"Message": fmt.Sprintf("Um erro ocorreu: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   "false",
			"Message": fmt.Sprintf("Inserido com sucesso! ID: %d", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
