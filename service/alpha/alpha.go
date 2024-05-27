package alpha

import (
	"context"
	"errors"
	"github.com/xoxoist/ut-tutor/external/alpha"
	"strings"
)

type (
	Service interface {
		GetProducts(ctx context.Context, request alpha.Request) (*alpha.Response, error)
	}
	service struct {
		alphaApi *alpha.API
	}
)

func (s *service) GetProducts(_ context.Context, request alpha.Request) (*alpha.Response, error) {
	data := s.alphaApi.GetAllProducts(request)
	if strings.Contains(data.Message, "ERROR_403") {
		return nil, errors.New("BAD_REQUEST")
	}
	if strings.Contains(data.Message, "ERROR_404") {
		return nil, errors.New("NOT_FOUND")
	}
	if strings.Contains(data.Message, "ERROR_401") {
		return nil, errors.New("AUTHORIZATION")
	}
	return &data, nil
}

func NewService(alphaApi *alpha.API) Service {
	return &service{alphaApi: alphaApi}
}
