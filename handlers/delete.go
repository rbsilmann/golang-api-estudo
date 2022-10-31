package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rbsilmann/api-estudo/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao realizar o parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Foram removidos mais registros do que o esperado: %d", rows)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "dado removido com sucesso!",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return
}
