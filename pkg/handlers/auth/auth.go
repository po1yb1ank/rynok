package auth

import (
	"context"
	"encoding/json"
	"github.com/po1yb1ank/rynok/pkg/logger"
	"github.com/po1yb1ank/rynok/tools"
	"github.com/sirupsen/logrus"
	"net/http"
)
var log *logrus.Logger

func LoginHandlerCtx(ctx context.Context, n http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log = logger.GetLogger(ctx)
		log.Println("login handler used")
		n.ServeHTTP(w, r)
	})
}
func RegisterHandlerCtx(ctx context.Context, n http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User

		log = logger.GetLogger(ctx)
		log.Println("register handler used")

		w.Header().Set("Content-Type","application/json")
		json.NewDecoder(r.Body).Decode(&user)
		user.Password = tools.GetHash(ctx,[]byte(user.Password))

	})
}
