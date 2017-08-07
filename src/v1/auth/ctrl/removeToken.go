package authCtrl

import (
	"net/http"
	"time"

	"github.com/Joematpal/test-api/src/v1/utils"
)

// RemoveToken from
func RemoveToken() utils.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			deleteCookie := http.Cookie{
				Name:    "Auth",
				Value:   "none",
				Expires: time.Now(),
			}

			http.SetCookie(w, &deleteCookie)
			h.ServeHTTP(w, r)
		})
	}
}
