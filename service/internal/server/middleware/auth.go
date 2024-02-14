package middleware

import (
	"net/http"

	"github.com/Seiya-Tagami/pecopeco-service/internal/presentation/responder"
	"github.com/Seiya-Tagami/pecopeco-service/internal/util/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rawIDToken := r.Header.Get("Authorization")
		_, err := jwt.Verify(ctx, rawIDToken)
		if err != nil {
			responder.ReturnStatusUnauthorized(w, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
