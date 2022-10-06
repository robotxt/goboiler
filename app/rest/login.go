package rest

import (
	"encoding/json"
	"net/http"
)

type loginData struct {
	Email    string
	Password string
}

type loginResponse struct {
	Token string
	Email string
}

func (srv RestService) LoginAPI(w http.ResponseWriter, r *http.Request) {
	ld := loginData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ld)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := srv.CoreSrv.BasicLoginAuth(ld.Email, ld.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srv.Log.Info("token: ", token)

	respondJSON(w, http.StatusOK, &loginResponse{
		Token: token.Token,
		Email: ld.Email,
	})
}
