package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/po1yb1ank/rynok/pkg/handlers/auth"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)
var log = logrus.New()
func main(){
	logrus.Info("Application started")

	router := mux.NewRouter()
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 15)
	http.Handle("/login",auth.LoginHandlerCtx(ctx, router))
	http.Handle("/register", auth.RegisterHandlerCtx(ctx,router))


}
