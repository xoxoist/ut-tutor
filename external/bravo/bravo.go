package bravo

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

func (b *API) GetAllStores(request Request) Response {
	if strings.Contains(request.Query.Value, "403") {
		return Response{
			Message: "ERROR_403",
			Status:  http.StatusBadRequest,
			Stores:  nil,
		}
	}
	if strings.Contains(request.Query.Value, "404") {
		return Response{
			Message: "ERROR_404",
			Status:  http.StatusNotFound,
			Stores:  nil,
		}
	}
	if strings.Contains(request.Query.Value, "401") {
		return Response{
			Message: "ERROR_401",
			Status:  http.StatusUnauthorized,
			Stores:  nil,
		}
	}

	var stores []Store
	for i := 0; i < request.Limit; i++ {
		stores = append(stores, Store{
			Name:    "Product Name for" + " " + strconv.Itoa(i),
			Address: "Address Name for " + " " + strconv.Itoa(i),
			Phone:   "081299882288",
		})
	}

	return Response{
		Message: "Success",
		Status:  200,
		Stores:  stores,
	}
}
