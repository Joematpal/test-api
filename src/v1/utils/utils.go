package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// RespondWithError responses with Error in the form of JSON.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON responses with JSON.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Adapter this.
type Adapter func(http.Handler) http.Handler

// Adapt this.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

// Logging is an example of an emplimentatin of an adapter.
func Logging() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}

// Validate checks if the user is logged in.
func Validate() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			cookie, err := req.Cookie("Auth")
			fmt.Println("validate", cookie, nil)
			if err != nil {
				res.Header().Set("Content-Type", "text/html")
				fmt.Fprint(res, "Unauthorized - Please login <br>")
				fmt.Fprintf(res, "<a href=\"login\"> Login </a>")
				return
			}

			// token, err := jwt.ParseWithClaims(cookie.Value, &Claims{},
			// 	func(token *jwt.Token) (interface{}, error) {
			// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// 			return nil, fmt.Errorf("Unexpected signing method")
			// 		}
			// 		return []byte("secret"), nil
			// 	},
			// )

			if err != nil {
				res.Header().Set("Content-Type", "text/html")
				fmt.Fprint(res, "Unauthorized - Please login <br>")
				fmt.Fprintf(res, "<a href=\"login\"> Login </a>")
				return
			}

			// if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			// 	ctx := context.WithValue(req.Context(), MyKey, *claims)
			// 	h(res, req.WithContext(ctx))
			// } else {
			// 	res.Header().Set("Content-Type", "text/html")
			// 	fmt.Fprint(res, "Unauthorized - Please login <br>")
			// 	fmt.Fprintf(res, "<a href=\"login\"> Login </a>")
			// 	return
			// }
		})
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
