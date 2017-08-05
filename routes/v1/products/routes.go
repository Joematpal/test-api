package products

import (
	"github.com/Joematpal/test-api/routes/v1/utils"
	"github.com/Joematpal/test-api/version"
)

// Routes for products
func Routes(v version.V1) {
	v.Subrouter.HandleFunc("/product/{id}", GetProduct(v.DB)).Methods("GET")
	v.Subrouter.HandleFunc("/product", CreateProduct(v.DB)).Methods("POST")
	v.Subrouter.HandleFunc("/product/{id}", UpdateProduct(v.DB)).Methods("PUT")
	v.Subrouter.HandleFunc("/product/{id}", DeleteProduct(v.DB)).Methods("DELETE")

	v.Subrouter.Handle("/products",
		utils.Adapt(
			GetProducts(v.DB),
			// utils.Validate(),
			utils.Logging(),
		)).Methods("GET")
}
