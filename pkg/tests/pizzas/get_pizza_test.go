package pizza_tests

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"pizza/pkg/common/models"
	"regexp"
	"testing"
)

func (s *PizzaTestSuite) TestApi_GetById() {
	var (
		id, _       = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
		url         = fmt.Sprintf("%s/%s", s.baseUri, id.String())
		request, _  = http.NewRequest(http.MethodGet, url, nil)
		recorder    = httptest.NewRecorder()
		response    models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), name, description))

	s.router.ServeHTTP(recorder, request)

	json.Unmarshal([]byte(recorder.Body.String()), &response)
	s.T().Run("should return status code 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
	s.T().Run("should return expected pizza", func(t *testing.T) {
		require.Nil(t, deep.Equal(&models.Pizza{ID: id, Name: name, Description: description}, &response))
	})
}

func (s *PizzaTestSuite) TestApi_GetById_NotFound() {
	var (
		id, _      = uuid.NewUUID()
		url        = fmt.Sprintf("%s/%s", s.baseUri, id.String())
		request, _ = http.NewRequest(http.MethodGet, url, nil)
		recorder   = httptest.NewRecorder()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows(nil))

	s.router.ServeHTTP(recorder, request)
	s.T().Run("should return status code 404", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func (s *PizzaTestSuite) TestApi_GetById_MalformedRequest() {
	var (
		id         = "this is not a uuid"
		url        = fmt.Sprintf("%s/%s", s.baseUri, id)
		request, _ = http.NewRequest(http.MethodGet, url, nil)
		recorder   = httptest.NewRecorder()
	)

	s.router.ServeHTTP(recorder, request)
	s.T().Run("should return status code 404", func(t *testing.T) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
