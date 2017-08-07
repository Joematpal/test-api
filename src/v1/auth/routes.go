package auth

import (
	"github.com/Joematpal/test-api/src/v1/auth/ctrl"
	"github.com/Joematpal/test-api/src/v1/utils"
	"github.com/Joematpal/test-api/src/v1/version"
)

// Routes for auth
func Routes(v version.V1) {
	v.Subrouter.Handle("/login",
		utils.Adapt(
			nil,
			authCtrl.SetToken(),
			authCtrl.CheckUser(v.DB),
		)).Methods("POST")

	v.Subrouter.Handle("/signup",
		utils.Adapt(
			nil,
			// SetToken(),
			authCtrl.SignUp(v.DB),
			authCtrl.PassConfirm(),
		)).Methods("POST")

	v.Subrouter.Handle("/logout",
		utils.Adapt(
			nil,
			authCtrl.RemoveToken(),
		)).Methods("GET")
}
