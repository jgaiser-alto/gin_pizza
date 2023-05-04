package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pizza/pkg/common/models"
	"pizza/pkg/pizzas"
	"regexp"
)

func (s *PizzaTestSuite) TestExpectedPizzaIsUpdated() {
	var (
		id, _ = uuid.NewUUID()
		url   = fmt.Sprintf("%s/%s", s.baseUri, id.String())
		body  = pizzas.UpdatePizzaRequestBody{
			Name:        "The Kramer",
			Description: "you make your own pie",
		}
		jsonBody, _   = json.Marshal(body)
		recorder      = httptest.NewRecorder()
		request, _    = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
		expectedPizza = &models.Pizza{ID: id, Name: body.Name, Description: body.Description}
		response      models.Pizza
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), "old name", "old body"))

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "pizzas"`)).
		WithArgs(body.Name, body.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), body.Name, body.Description))
	s.mock.ExpectCommit()

	s.router.ServeHTTP(recorder, request)
	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusOK, recorder.Code)
	if diff := deep.Equal(expectedPizza, &response); diff != nil {
		s.T().Error(diff)
	}
}

func (s *PizzaTestSuite) TestMalformedUpdateRequest() {
	var (
		id, _ = uuid.NewUUID()
		url   = fmt.Sprintf("%s/%s", s.baseUri, id.String())
		body  = pizzas.UpdatePizzaRequestBody{
			Name:        "NO Description",
			Description: "",
		}
		jsonBody, _ = json.Marshal(body)
		recorder    = httptest.NewRecorder()
		request, _  = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	)

	s.router.ServeHTTP(recorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, recorder.Code)
}

func (s *PizzaTestSuite) TestUpdateExceptionIsThrown() {
	var (
		id, _ = uuid.NewUUID()
		url   = fmt.Sprintf("%s/%s", s.baseUri, id.String())
		body  = pizzas.UpdatePizzaRequestBody{
			Name:        "new pizza",
			Description: "a new pizza pie",
		}
		jsonBody, _ = json.Marshal(body)
		recorder    = httptest.NewRecorder()
		request, _  = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), "old name", "old body"))

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`UPDATE "pizzas"`)).
		WithArgs(body.Name, body.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("something didn't work"))
	s.mock.ExpectRollback()

	s.router.ServeHTTP(recorder, request)

	assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)
}
