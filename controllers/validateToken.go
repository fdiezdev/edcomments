package controllers

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/fdiezdev/edcomments/commons"
	"github.com/fdiezdev/edcomments/models"
)

type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

// ValidateToken validates user (client) token
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message

	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)

	if err != nil {
		m.StatusCode = http.StatusUnauthorized

		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token ha expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del token no coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "El token no es válido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}

	if token.Valid {
		ctx := context.WithValue(r.Context(), userCtxKey, token.Claims.(*models.Claim).User)

		next(w, r.WithContext(ctx))
	} else {
		m.StatusCode = http.StatusUnauthorized
		m.Message = "Error: Su token no es válido"
		commons.DisplayMessage(w, m)
	}
}
