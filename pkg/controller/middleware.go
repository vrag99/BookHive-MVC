package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/golang-jwt/jwt/v4"
)

var JWTContextKey types.Claims

func GenerateJWT(userData types.UserData) string {
	config, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(config.AccessTokenSecret)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = userData.Id
	claims["username"] = userData.Username
	claims["admin"] = userData.Admin
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	return tokenStr
}

func DecodeJWT(tokenStr string) (jwt.MapClaims, error) {
	config, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
	secretKey := []byte(config.AccessTokenSecret)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

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
			claims, err := DecodeJWT(token)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				customClaims := types.Claims {
					Id: claims["id"],
					Username: claims["username"].(string),
					IsAdmin: claims["admin"] != 0.0,
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
