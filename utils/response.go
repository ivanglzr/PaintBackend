package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ivanglzr/PaintBackend/models"
)

func JSONResponse(w http.ResponseWriter, statusCode int, message string) {
	var res models.Response
	res.StatusCode = statusCode
	res.Message = message

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
