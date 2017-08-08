package authCtrl

import (
	"database/sql"
	"net/http"

	"github.com/Joematpal/test-golang-api/src/v1/users"
	"github.com/Joematpal/test-golang-api/src/v1/utils"
	"github.com/Joematpal/test-golang-api/src/v1/utils/respond"
	uuid "github.com/satori/go.uuid"
)

// SignUp creates a newUser
func SignUp(db *sql.DB) utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer db.Close()
			u := users.User{}

			if err := utils.Decode(r, &u); err != nil {
				respond.With(w, r, http.StatusBadRequest, nil, "err at decode")
				return
			}

			hash, err := utils.HashPassword(u.Password)
			if err != nil {
				//TODO: make sure that the front end and the backend are communicating properly on duplicate information.
				respond.With(w, r, http.StatusPartialContent, nil, err)
				return
			}
			u.PasswordHash = hash
			u.ID = uuid.NewV4()
			if err := u.CreateUser(db); err != nil {
				respond.With(w, r, http.StatusPartialContent, nil, err)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
