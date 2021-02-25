package tools

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

func GetHash(ctx context.Context,pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		ctx.Value().Println(err)
	}
	return string(hash)
}

