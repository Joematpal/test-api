package authCtrl

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Joematpal/test-api/src/v1/auth/mdl"
	"github.com/Joematpal/test-api/src/v1/users"

	"github.com/Joematpal/test-api/src/v1/utils"
	"github.com/Joematpal/test-api/src/v1/utils/ctxKeys"
	"github.com/Joematpal/test-api/src/v1/utils/respond"
	jwt "github.com/dgrijalva/jwt-go"
)

// Validate checks if the token on the request has a uuid that matches the db.
func Validate(db *sql.DB) utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("Auth")
			if err != nil {
				respond.With(w, r, http.StatusBadRequest, nil, err)
				return
			}

			token, e := jwt.ParseWithClaims(cookie.Value, &authMdl.Session{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method")
					}
					return []byte(os.Getenv("TOKEN_SECRET")), nil
				})
			if e != nil {
				respond.With(w, r, http.StatusBadRequest, nil, e)
				return
			}

			claims, ok := token.Claims.(*authMdl.Session)
			if !ok && !token.Valid {

				respond.With(w, r, http.StatusBadRequest, nil, "please provide a token")
				return
			}

			u := users.User{ID: claims.ID}

			if err := u.CheckID(db); err != nil {
				respond.With(w, r, http.StatusBadRequest, nil, err)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeys.Auth("userid"), u)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
