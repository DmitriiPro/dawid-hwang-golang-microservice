package handler

import (
	"context"
	"davidHwang/ecomm/token"
	"fmt"
	"net/http"
	"strings"
)

type authKey struct {
}

// * middleware администратора
func GetAdminMiddlewareFunc(tokenMaker *token.JWTMaker) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//* сначала прочитаем заголовок авторизации
			//* проверим токен на валидность
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)

			if err != nil {
				http.Error(w, fmt.Sprintf("error verifying token: %w", err), http.StatusUnauthorized)
				return
			}

			//! проверка на то что является ли пользователь администратором
			if !claims.IsAdmin {
				http.Error(w, "user is not admin", http.StatusForbidden)
				return
			}


			//* передадим в контекст запроса токен и пользователя
			ctx := context.WithValue(r.Context(), authKey{}, claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}
func GetAuthMiddlewareFunc(tokenMaker *token.JWTMaker) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			//* сначала прочитаем заголовок авторизации
			//* проверим токен на валидность
			claims, err := verifyClaimsFromAuthHeader(r, tokenMaker)

			if err != nil {
				http.Error(w, fmt.Sprintf("error verifying token: %w", err), http.StatusUnauthorized)
				return
			}
			//* передадим в контекст запроса токен и пользователя
			ctx := context.WithValue(r.Context(), authKey{}, claims)

			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}

// * вспомогательная функция
func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker *token.JWTMaker) (*token.UserClaims, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)

	if len(fields) != 2 || fields[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header")
	}

	token := fields[1]
	claims, err := tokenMaker.VerifyToken(token)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil

}
