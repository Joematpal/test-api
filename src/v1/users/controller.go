package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joematpal/test-api/src/v1/utils/respond"
	"github.com/gorilla/mux"
)

// GetUser HandlerFunc
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		u := User{Username: username}

		if err := u.getUser(db); err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, u, nil)
	}
}

// GetUsers does this.
func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		u := User{}
		fmt.Println("we are in users/controller")
		users, err := u.getUsers(db, start, count)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, users, nil)
	}
}
