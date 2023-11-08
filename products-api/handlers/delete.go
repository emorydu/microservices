package handlers

import (
	"errors"
	"github.com/emorydu/microservices/products-api/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	204: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id:", id)

	err := data.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, data.ErrProductNotFound) {
			p.l.Println("[ERROR] deleting record id does not exist")

			w.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
