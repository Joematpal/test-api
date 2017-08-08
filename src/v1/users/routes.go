package users

import (
	"github.com/Joematpal/test-golang-api/src/v1/utils"
	"github.com/Joematpal/test-golang-api/src/v1/version"
)

// Routes for users
func Routes(v version.V1) {
	v.Subrouter.Handle("/users",
		utils.Adapt(
			GetUsers(v.DB),
		)).Methods("GET")

	v.Subrouter.Handle("/user/{id}",
		utils.Adapt(
			GetUser(v.DB),
		)).Methods("GET")
}
