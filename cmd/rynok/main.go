package main

import (
	"github.com/gorilla/mux"
	"github.com/po1yb1ank/rynok/pkg/auth"
	"github.com/sirupsen/logrus"
)

func main(){
	logrus.Info("Application started")

	router := mux.NewRouter()
	//ctx, _ := context.WithTimeout(context.Background(), time.Second * 15)

	router.HandleFunc("/login",auth.LoginHandler)
	router.HandleFunc("/register", auth.RegisterHandler)

}
