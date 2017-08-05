package users

import (
	"github.com/Joematpal/test-api/routes/v1/utils"
	"github.com/Joematpal/test-api/version"
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
