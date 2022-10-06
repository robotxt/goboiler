package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type UserCtxKey string

type AuthUser struct {
	UID   string
	Email string
}

func (srv RestService) SecureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("HTTP_AUTHORIZATION")

		json.NewEncoder(w).Encode(r)
		token := strings.TrimSpace(header)

		if token == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing HTTP_AUTHORIZATION Header")
			return
		}

		user, err := srv.CoreSrv.ValidateUserToken(token)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Forbidden")
			return
		}

		ctx := context.WithValue(context.Background(), UserCtxKey("AuthUser"), AuthUser{
			UID:   user.UID,
			Email: user.Email,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
