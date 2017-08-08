package authCtrl

import (
	"net/http"
	"os"
	"time"

	authMdl "github.com/Joematpal/test-golang-api/src/v1/auth/mdl"
	"github.com/Joematpal/test-golang-api/src/v1/users"
	"github.com/Joematpal/test-golang-api/src/v1/utils"
	"github.com/Joematpal/test-golang-api/src/v1/utils/ctxKeys"
	"github.com/Joematpal/test-golang-api/src/v1/utils/respond"
	jwt "github.com/dgrijalva/jwt-go"
)

// SetToken is middleware that sets the user's token.
func SetToken() utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uCtx := r.Context().Value(ctxKeys.Auth("userid"))
			u := uCtx.(users.User)

			expireToken := time.Now().Add(time.Hour * 8).Unix()
			expireCookie := time.Now().Add(time.Hour * 8)

			claims := authMdl.Session{
				u.Username,
				u.ID,
				u.Email,
				jwt.StandardClaims{
					ExpiresAt: expireToken,
					Issuer:    os.Getenv("SERVER_DOMAIN"),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			signedToken, _ := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

			cookie := http.Cookie{
				Name:     "Auth",
				Value:    signedToken,
				Expires:  expireCookie, // 30 min
				HttpOnly: true,
				Path:     "/",
				Domain:   os.Getenv("SERVER_DOMAIN"),
				Secure:   true,
			}

			http.SetCookie(w, &cookie)

			respond.With(w, r, http.StatusOK, u.Public(), nil)
			h.ServeHTTP(w, r)
		})
	}
}
