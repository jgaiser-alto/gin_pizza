package tests

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
)

func (s *PizzaTestSuite) TestExpectedPizzaIsReturned() {
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
	assert.Equal(s.T(), http.StatusOK, recorder.Code)
	require.Nil(s.T(), deep.Equal(&models.Pizza{ID: id, Name: name, Description: description}, &response))
}

func (s *PizzaTestSuite) TestPizzaNotFound() {
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
	assert.Equal(s.T(), http.StatusNotFound, recorder.Code)
}

func (s *PizzaTestSuite) TestMalformedRequest() {
	var (
		id         = "this is not a uuid"
		url        = fmt.Sprintf("%s/%s", s.baseUri, id)
		request, _ = http.NewRequest(http.MethodGet, url, nil)
		recorder   = httptest.NewRecorder()
	)

	s.router.ServeHTTP(recorder, request)
	assert.Equal(s.T(), http.StatusNotFound, recorder.Code)
}
