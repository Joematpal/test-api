package v1

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Joematpal/test-api/routes/v1/auth"
	"github.com/Joematpal/test-api/routes/v1/products"
	"github.com/Joematpal/test-api/routes/v1/users"
	"github.com/Joematpal/test-api/version"
	"github.com/gorilla/mux"
	// postgres
	_ "github.com/lib/pq"
)

// V1 routes is "/api/v1"
// type V1 struct {
// 	DB        *sql.DB
// 	Subrouter *mux.Router
// }

// Initialize this thing
func Initialize(user string, password string, dbname string, newRouter *mux.Router) {
	// sslmode=disable this stuff below if wrong
	var v1 = version.V1{}
	certPath := "cert/server.pem"
	keyPath := "cert/server.key"

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable sslcert=%s sslkey=%s", user, password, dbname, certPath, keyPath)

	var err error
	v1.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	v1.Subrouter = newRouter.PathPrefix("/api/v1").Subrouter()

	// Initialize Routes
	auth.Routes(v1)
	users.Routes(v1)
	products.Routes(v1)
}
