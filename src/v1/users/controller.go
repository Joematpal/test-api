package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joematpal/test-api/src/v1/utils"
	"github.com/gorilla/mux"
)

// GetUser HandlerFunc
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		u := User{ID: id}

		if err := u.getUser(db); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, u)
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
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, users)
	}
}
