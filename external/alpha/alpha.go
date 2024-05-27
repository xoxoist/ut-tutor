package alpha

import (
	"github.com/stretchr/testify/mock"
	"net/http"
	"strconv"
	"strings"
)

type API struct {
	mock.Mock
}

func NewAPI() *API {
	return &API{}
}

func (a *API) GetAllProducts(request Request) Response {
	if strings.Contains(request.Query.Value, "403") {
		return Response{
			Message:  "ERROR_403",
			Status:   http.StatusBadRequest,
			Products: nil,
		}
	}
	if strings.Contains(request.Query.Value, "404") {
		return Response{
			Message:  "ERROR_404",
			Status:   http.StatusNotFound,
			Products: nil,
		}
	}
	if strings.Contains(request.Query.Value, "401") {
		return Response{
			Message:  "ERROR_401",
			Status:   http.StatusUnauthorized,
			Products: nil,
		}
	}

	var products []Product
	for i := 0; i < request.Limit; i++ {
		products = append(products, Product{
			Name:     "Product Name for" + " " + strconv.Itoa(i),
			Category: "Category Name for " + " " + strconv.Itoa(i),
			Price:    "20.000",
		})
	}

	return Response{
		Message:  "Success",
		Status:   200,
		Products: products,
	}
}
