package api

import (
	"encoding/json"
	"net/http"
	"yandexGoCalc/internal/orchestrator"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	ID string `json:"id"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request, o *orchestrator.Orchestrator) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Bad request"}`, http.StatusBadRequest)
		return
	}

	id := o.AddExpression(req.Expression)

	response := Response{ID: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
