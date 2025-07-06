// internal/api/handlers.go
package api

import (
	"encoding/json"
	"net/http"

	"featurn/internal/eval"
	"featurn/internal/flags"
)

type EvalRequest struct {
	FlagKey string                 `json:"flag_key"`
	Context map[string]interface{} `json:"context"`
}

func EvaluateFlagHandler(w http.ResponseWriter, r *http.Request) {
	var req EvalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	flag, found := flags.GetFlagByKey(req.FlagKey)
	if !found {
		http.Error(w, "Flag not found", http.StatusNotFound)
		return
	}

	result := eval.Evaluate(flag, req.Context)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"value": result,
	})
}
