package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Joematpal/test-api/routes/v1/users"
	"github.com/Joematpal/test-api/routes/v1/utils"
)

// SetToken is middleware that sets the user's token.
func SetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// CheckUser is the middleware
func CheckUser(db *sql.DB) utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := users.User{}

			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
				return
			}

			// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			// if err != nil {
			// 	panic(err)
			// }
			// u.PasswordHash = string(hashedPassword)
			// if err := u.UpdateUser(db); err != nil {
			//
			// }

			if err := u.CheckPass(db); err != nil {
				fmt.Println("user does not exist")
			}

			fmt.Println(u.PasswordHash)
			match := utils.CheckPasswordHash(u.Password, u.PasswordHash)
			fmt.Println("Match:   ", match)
		})
	}
}

// SignUp creates a newUser
func SignUp(db *sql.DB) utils.Adapter {
	return func(http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u := users.User{}
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
				return
			}

			if hash, err := utils.HashPassword(u.Password); err != nil {

			} else {
				u.PasswordHash = hash
			}
			// I need to create a uuid
			fmt.Print(u)
			if err := u.CreateUser(db); err != nil {
				fmt.Println(
					"err", err.Error(),
				)
			}
		})
	}
}

// PassConfirm checks the password and passwordConfirm.
func PassConfirm() utils.Adapter {
	return func(http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}

// RemoveToken from
func RemoveToken() utils.Adapter {
	return func(http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}
