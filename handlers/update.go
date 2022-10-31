package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rbsilmann/api-estudo/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao realizar o parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var todo models.ToDo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao realizar o decode do JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Foram atualizados mais registros do que o esperado: %d", rows)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso!",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return
}
