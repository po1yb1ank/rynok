package auth

import (
	"context"
	"encoding/json"
	"github.com/po1yb1ank/rynok/tools"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){


}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var user User

	w.Header().Set("Content-Type","application/json")
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = tools.GetHash(context.Background(),[]byte(user.Password))



}
