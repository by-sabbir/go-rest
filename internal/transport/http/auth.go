package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		fmt.Println("header parts: ", authHeaderParts[1])
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
		if validateToken(authHeaderParts[1]) {
			original(w, r)
		} else {
			http.Error(w, "not authorized", http.StatusForbidden)
			return
		}
	}
}

func validateToken(tokenString string) bool {
	secret := []byte("golang")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	fmt.Println("token extract: ", token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token claims: ", claims)
		return true
	} else {
		fmt.Println(err)
	}
	return false
}
