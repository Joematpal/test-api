package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Joematpal/test-api/routes/v1"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	router := mux.NewRouter()
	// V1 := v1.V1{}

	v1.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		router,
	)

	run(":8080", router)
}

func run(addr string, route *mux.Router) {
	certPath := "cert/server.pem"
	keyPath := "cert/server.key"
	log.Fatal(http.ListenAndServeTLS(addr, certPath, keyPath, route))
}
