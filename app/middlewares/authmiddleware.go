package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
	"time"
)

type AuthMiddleware struct {
	accessSecret string
	superUser    string
}

func NewAuthMiddleware(accessSecret, superUser string) *AuthMiddleware {
	return &AuthMiddleware{
		accessSecret: accessSecret,
		superUser:    superUser,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			jwtToken *jwt.Token
			err      error
		)

		jwtToken, err = token.NewTokenParser().ParseToken(r, m.accessSecret, "")

		if err != nil || !jwtToken.Valid {
			http.Error(w, "the token is invalid", http.StatusUnauthorized)
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)
		if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
			http.Error(w, "the token is expired", http.StatusUnauthorized)
			return
		}

		fmt.Println("8888888888888888")
		next(w, r)
	}
}
