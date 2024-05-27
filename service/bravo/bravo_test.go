package bravo

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/xoxoist/ut-tutor/external"
	"github.com/xoxoist/ut-tutor/external/bravo"
	"github.com/xoxoist/ut-tutor/shared"
	"go.uber.org/dig"
	"net/http"
	"testing"
)

// testAlphaServiceTestSuite embedding suite.Suite of testify
type testAlphaServiceTestSuite struct {
	suite.Suite
	di *dig.Container
}

// TestAlphaServiceTestSuite initialize test suite
func TestAlphaServiceTestSuite(t *testing.T) {
	suite.Run(t, new(testAlphaServiceTestSuite))
}

// SetupTest setup dependencies injection to used for testing purpose
func (s *testAlphaServiceTestSuite) SetupTest() {
	s.di = dig.New()
	err := s.di.Provide(shared.NewEnvs)
	err = external.RegisterDependencies(s.di)
	err = s.di.Provide(NewService)
	s.Assert().NoError(err)
}

// GetProducts testing get all product
func (s *testAlphaServiceTestSuite) TestGetProducts() {
	s.Assert().NoError(s.di.Invoke(func(
		envs *shared.Envs,
		externals external.DependenciesHolder,
		bravoSvc Service,
	) {
		// make test cases
		tests := []struct {
			TestCaseName      string
			Request           bravo.Request
			MockedAPIResponse bravo.Response
			WantError         bool
			ErrorString       string
		}{
			{
				TestCaseName: "Test 403 Bad Request",
				Request: bravo.Request{
					Query: bravo.Query{
						Keyword: "Testing",
						Value:   "403",
					},
					Limit: 1,
					Page:  0,
				},
				WantError:   true,
				ErrorString: "BAD_REQUEST",
				MockedAPIResponse: bravo.Response{
					Message: "ERROR_403",
					Status:  http.StatusBadRequest,
					Stores:  nil,
				},
			},
			{
				TestCaseName: "Test 404 Not Found",
				Request: bravo.Request{
					Query: bravo.Query{
						Keyword: "Testing",
						Value:   "404",
					},
					Limit: 1,
					Page:  0,
				},
				WantError:   true,
				ErrorString: "NOT_FOUND",
				MockedAPIResponse: bravo.Response{
					Message: "ERROR_404",
					Status:  http.StatusNotFound,
					Stores:  nil,
				},
			},
			{
				TestCaseName: "Test 401 Authorization",
				Request: bravo.Request{
					Query: bravo.Query{
						Keyword: "Testing",
						Value:   "401",
					},
					Limit: 1,
					Page:  0,
				},
				WantError:   true,
				ErrorString: "AUTHORIZATION",
				MockedAPIResponse: bravo.Response{
					Message: "ERROR_401",
					Status:  http.StatusUnauthorized,
					Stores:  nil,
				},
			},
		}

		// scan test cases
		for _, test := range tests {
			s.Run(test.TestCaseName, func() {
				// setup test expectation
				externals.Bravo.
					On("GetAllProducts", mock.Anything, test.Request).
					Return(test.MockedAPIResponse)

				// start test based on test case
				_, err := bravoSvc.GetStores(context.Background(), test.Request)
				if test.WantError {
					s.Assert().Error(err)
					s.Assert().Equal(err.Error(), test.ErrorString)
				} else {
					s.Assert().NoError(err)
				}
			})
		}
	}))
}
