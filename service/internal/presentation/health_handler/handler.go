package healthhandler

import (
	"net/http"

	"github.com/Seiya-Tagami/pecopeco-service/internal/db"
	"github.com/Seiya-Tagami/pecopeco-service/internal/presentation/responder"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	db, err := db.NewMySQL()
	if err != nil {
		res := HealthResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed to get connection database",
			Detail:  err.Error(),
		}
		responder.RespondJson(w, res, res.Status)
		return
	}
	defer db.Close()

	res := HealthResponse{
		Status:  http.StatusOK,
		Message: "success to connect database",
	}
	responder.RespondJson(w, res, res.Status)
}
