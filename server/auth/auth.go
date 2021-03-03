package auth

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strings"
)

var tkn struct {
	Jwt string
}

// ValidateToken - Check token is valid
func ValidateToken(r *http.Request) (payload string, err error) {
	h := r.Header
	lines := h["Cookie"]
	session := lines[0]
	keyValue := strings.Split(session, "express:sess=")

	sessionValue := keyValue[1]
	decoded, err := jwt.DecodeSegment(sessionValue)
	json.Unmarshal(decoded, &tkn)

	token, err := jwt.Parse(tkn.Jwt, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	var p string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		p = getPayload(claims)
	}

	return p, nil
}

func getPayload(data interface{}) string {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(b)
}
