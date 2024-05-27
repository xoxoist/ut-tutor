package bravo

import (
	"context"
	"errors"
	"github.com/xoxoist/ut-tutor/external/bravo"
	"strings"
)

type (
	Service interface {
		GetStores(ctx context.Context, request bravo.Request) (*bravo.Response, error)
	}
	service struct {
		bravoApi *bravo.API
	}
)

func (s *service) GetStores(_ context.Context, request bravo.Request) (*bravo.Response, error) {
	data := s.bravoApi.GetAllStores(request)
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

func NewService(bravoApi *bravo.API) Service {
	return &service{bravoApi: bravoApi}
}
