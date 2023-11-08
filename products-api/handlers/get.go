package handlers

import (
	"errors"
	"github.com/emorydu/microservices/products-api/data"
	"net/http"
)

// swagger:route GET /products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, w)
	if err != nil {
		// we should never be here but log the error just encase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch {
	case err == nil:
	case errors.Is(err, data.ErrProductNotFound):
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, w)
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		// we should never be here but log the error just encase
		p.l.Println("[ERROR] serializing product", err)
	}
}
