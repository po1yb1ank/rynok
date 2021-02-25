package tools

import (
	"context"
	"github.com/po1yb1ank/rynok/pkg/logger"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)
var log *logrus.Logger
func GetHash(ctx context.Context,pwd []byte) string {
	log = logger.GetLogger(ctx)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

