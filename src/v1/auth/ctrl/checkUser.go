package authCtrl

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Joematpal/test-golang-api/src/v1/users"
	"github.com/Joematpal/test-golang-api/src/v1/utils"
	"github.com/Joematpal/test-golang-api/src/v1/utils/ctxKeys"
	"github.com/Joematpal/test-golang-api/src/v1/utils/respond"
)

// CheckUser is the middleware
func CheckUser(db *sql.DB) utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			u := users.User{}
			if err := utils.Decode(r, &u); err != nil {
				respond.With(w, r, http.StatusBadRequest, nil, "err at decode")
				return
			}

			if err := u.CheckPass(db); err != nil {
				respond.With(w, r, http.StatusBadRequest, "what", "The Username or Password provided could not be authenticated.")
				return
			}

			match := utils.CheckPasswordHash(u.Password, u.PasswordHash)
			if !match {
				respond.With(w, r, http.StatusBadRequest, nil, "Wrong password")
				return
			}

			// fmt.Println(r.Context().Value(ctxKeys.Auth("userid")))
			ctx := context.WithValue(r.Context(), ctxKeys.Auth("userid"), u)

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
