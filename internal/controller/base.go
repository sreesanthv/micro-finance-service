package controller

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
}

func (c *Controller) Decode(r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(dest)
}

func (c *Controller) SendMessage(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func (c *Controller) SendResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonResp, _ := json.Marshal(body)
	w.Write(jsonResp)
}

func (a *AccountController) hash(val string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
	return string(hash), err
}

func (a *AccountController) compareHash(hash string, val string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(val))
	return err == nil
}
