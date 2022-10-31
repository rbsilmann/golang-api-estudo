package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rbsilmann/api-estudo/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.SelectAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao realizar o parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var todo models.ToDo
	todo, err = models.SelectOne(int64(id))
	if err != nil {
		log.Printf("Erro ao recuperar o registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
	return
}
