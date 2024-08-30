package healthhandler

import (
	"net/http"

	"github.com/taga3s/pecopeco-service/internal/db"
	"github.com/taga3s/pecopeco-service/internal/presentation/responder"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := db.CheckConnection()
	if err != nil {
		responder.ReturnStatusInternalServerError(w, err)
		return
	}
	res := HealthResponse{
		Status:  http.StatusOK,
		Message: "success to connect database",
	}
	responder.ReturnStatusOK(w, res)
}
