package authCtrl

import (
	"net/http"

	"github.com/Joematpal/test-api/src/v1/users"
	"github.com/Joematpal/test-api/src/v1/utils"
	"github.com/Joematpal/test-api/src/v1/utils/respond"
)

// PassConfirm checks the password and passwordConfirm.
func PassConfirm() utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := users.NewUser{}

			if err := utils.Decode(r, &u); err != nil {
				respond.With(w, r, http.StatusBadRequest, nil, err)
				return
			}
			if u.Password != u.PasswordConfirm {
				respond.With(w, r, http.StatusPartialContent, nil, "password and password confirm do not match")
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}
