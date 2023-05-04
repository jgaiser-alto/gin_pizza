package tests

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
)

func (s *PizzaTestSuite) TestExpectedPizzasAreReturned() {
	var (
		id1, _      = uuid.NewUUID()
		id2, _      = uuid.NewUUID()
		id3, _      = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
		request, _  = http.NewRequest(http.MethodGet, s.baseUri, nil)
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

	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusOK, recorder.Code)
	assert.Equal(s.T(), 3, len(response))
}

func (s *PizzaTestSuite) TestEmptyResponse() {
	var (
		request, _ = http.NewRequest(http.MethodGet, s.baseUri, nil)
		recorder   = httptest.NewRecorder()
		response   []models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas"`)).
		WillReturnRows(sqlmock.NewRows(nil))

	s.router.ServeHTTP(recorder, request)

	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusOK, recorder.Code)
	assert.Equal(s.T(), 0, len(response))
}

func (s *PizzaTestSuite) TestGetExceptionIsThrown() {
	var (
		request, _ = http.NewRequest(http.MethodGet, s.baseUri, nil)
		recorder   = httptest.NewRecorder()
		response   []models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas"`)).
		WillReturnError(errors.New("something didn't work"))

	s.router.ServeHTTP(recorder, request)

	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusNotFound, recorder.Code)
}
