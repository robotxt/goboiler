package rest

import (
	"encoding/json"
	"net/http"
)

type basicRegistrationData struct {
	Email    string
	Password string
}

type basicRegistrationResp struct {
	Message string
	UID     string
}

func (srv RestService) registrationAPI(w http.ResponseWriter, r *http.Request) {
	rd := basicRegistrationData{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := srv.CoreSrv.BasicUserRegistration(rd.Email, rd.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, &basicRegistrationResp{
		Message: "Successfully Created.",
		UID:     user.UID,
	})
}
