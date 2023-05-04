package tests

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (s *PizzaTestSuite) TestExpectedPizzaIsCreated() {
	var (
		id, _ = uuid.NewUUID()
		body  = pizzas.AddPizzaRequestBody{
			Name:        "XXXL",
			Description: "BIG PIZZA FOR THE FAM",
		}
		jsonBody, _   = json.Marshal(body)
		recorder      = httptest.NewRecorder()
		request, _    = http.NewRequest(http.MethodPost, s.baseUri, bytes.NewBuffer(jsonBody))
		expectedPizza = &models.Pizza{ID: id, Name: body.Name, Description: body.Description}
		response      models.Pizza
	)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "pizzas"`)).
		WithArgs(body.Name, body.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), body.Name, body.Description))
	s.mock.ExpectCommit()

	s.router.ServeHTTP(recorder, request)
	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusCreated, recorder.Code)
	if diff := deep.Equal(expectedPizza, &response); diff != nil {
		s.T().Error(diff)
	}
}

func (s *PizzaTestSuite) TestMalformedCreateRequest() {
	var (
		body = pizzas.AddPizzaRequestBody{
			Name:        "NO Description",
			Description: "",
		}
		jsonBody, _ = json.Marshal(body)
		recorder    = httptest.NewRecorder()
		request, _  = http.NewRequest(http.MethodPost, s.baseUri, bytes.NewBuffer(jsonBody))
	)

	s.router.ServeHTTP(recorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, recorder.Code)
}

func (s *PizzaTestSuite) TestCreateExceptionIsThrown() {
	var (
		body = pizzas.AddPizzaRequestBody{
			Name:        "new pizza",
			Description: "a new pizza pie",
		}
		jsonBody, _ = json.Marshal(body)
		recorder    = httptest.NewRecorder()
		request, _  = http.NewRequest(http.MethodPost, s.baseUri, bytes.NewBuffer(jsonBody))
	)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "pizzas"`)).
		WithArgs(body.Name, body.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("something didn't work"))
	s.mock.ExpectRollback()

	s.router.ServeHTTP(recorder, request)

	assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)
}
