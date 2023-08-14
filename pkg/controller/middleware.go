package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"context"
	"log"
	"net/http"
	"strings"

	"golang.org/x/exp/slices"
)

var JWTContextKey types.Claims

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goThroughUrls := []string{"/", "/register", "/login", "/403", "/500"}
		if slices.Contains(goThroughUrls, r.URL.Path) || strings.Split(r.URL.Path, "/")[1] == "static" {
			next.ServeHTTP(w, r)
			return
		}
		cookie, err := r.Cookie("access-token")
		if err == nil {
			token := cookie.Value
			claims, err := utils.DecodeJWT(token)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				customClaims := types.Claims{
					Id:       claims["id"],
					Username: claims["username"].(string),
					IsAdmin:  claims["admin"] != 0.0,
				}
				ctx := context.WithValue(r.Context(), JWTContextKey, customClaims)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})
}

func CheckForAdmin(isAdminAuth bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims := r.Context().Value(JWTContextKey).(types.Claims)
			isAdmin := claims.IsAdmin
			isAdminDb, err := models.CheckAdmin(claims.Id)

			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
				return
			}
			if isAdmin == isAdminAuth && isAdmin == isAdminDb {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/forbiddenRequest", http.StatusSeeOther)
				return
			}
		})
	}
}
