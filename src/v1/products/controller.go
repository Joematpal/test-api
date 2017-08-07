package products

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Joematpal/test-api/src/v1/utils/respond"
	"github.com/gorilla/mux"
)

// GetProducts from ./model.go
func GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		p := Product{}

		products, err := p.getProducts(db, start, count)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			return
		}

		respond.With(w, r, http.StatusOK, nil, products)
	}
}

// GetProduct from ./model.go
func GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, nil, "Invalid product ID")
			return
		}

		p := Product{ID: id}
		if err := p.getProduct(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				respond.With(w, r, http.StatusNotFound, nil, "Product not found")
			default:
				respond.With(w, r, http.StatusInternalServerError, nil, err.Error())
			}
			return
		}

		respond.With(w, r, http.StatusOK, p, nil)
	}
}

// CreateProduct from ./model.go
func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// UpdateProduct from ./model.go
func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// DeleteProduct from ./model.go
func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
