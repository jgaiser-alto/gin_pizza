package pizzatests

import (
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pizza/pkg/common/models"
	"regexp"
	"testing"
)

func (s *PizzaTestSuite) TestApi_GetAll() {
	var (
		id1, _      = uuid.NewUUID()
		id2, _      = uuid.NewUUID()
		id3, _      = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
		request, _  = http.NewRequest(http.MethodGet, s.baseURI, nil)
		recorder    = httptest.NewRecorder()
		response    []models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id1, name, description).
			AddRow(id2, name, description).
			AddRow(id3, name, description),
		)

	s.router.ServeHTTP(recorder, request)

	json.Unmarshal(recorder.Body.Bytes(), &response)

	s.T().Run("should return status code 200", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
	s.T().Run("should return expected number of pizzas", func(t *testing.T) {
		assert.Equal(t, 3, len(response))
	})
}

func (s *PizzaTestSuite) TestApi_GetAll_EmptyCollection() {
	var (
		request, _ = http.NewRequest(http.MethodGet, s.baseURI, nil)
		recorder   = httptest.NewRecorder()
		response   []models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas"`)).
		WillReturnRows(sqlmock.NewRows(nil))

	s.router.ServeHTTP(recorder, request)

	json.Unmarshal(recorder.Body.Bytes(), &response)

	s.T().Run("should return status code 200", func(t *testing.T) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
	s.T().Run("should return expected number of pizzas", func(t *testing.T) {
		assert.Equal(t, 0, len(response))
	})
}

func (s *PizzaTestSuite) TestApi_GetAll_InternalServerError() {
	var (
		request, _ = http.NewRequest(http.MethodGet, s.baseURI, nil)
		recorder   = httptest.NewRecorder()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas"`)).
		WillReturnError(errors.New("something didn't work"))

	s.router.ServeHTTP(recorder, request)

	s.T().Run("should return status code 500", func(t *testing.T) {
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}
