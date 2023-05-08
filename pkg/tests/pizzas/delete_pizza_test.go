package pizzatests

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func (s *PizzaTestSuite) TestApi_DeleteById() {
	var (
		id, _       = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
		url         = fmt.Sprintf("%s/%s", s.baseURI, id.String())
		request, _  = http.NewRequest(http.MethodDelete, url, nil)
		recorder    = httptest.NewRecorder()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), name, description))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	s.router.ServeHTTP(recorder, request)
	s.T().Run("should return status code 200", func(t *testing.T) {
		assert.Equal(s.T(), http.StatusOK, recorder.Code)
	})
}

func (s *PizzaTestSuite) TestApi_DeleteById_NotFound() {
	var (
		id, _      = uuid.NewUUID()
		url        = fmt.Sprintf("%s/%s", s.baseURI, id.String())
		request, _ = http.NewRequest(http.MethodDelete, url, nil)
		recorder   = httptest.NewRecorder()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows(nil))

	s.router.ServeHTTP(recorder, request)
	s.T().Run("should return status code 404", func(t *testing.T) {
		assert.Equal(s.T(), http.StatusNotFound, recorder.Code)
	})
}

func (s *PizzaTestSuite) TestApi_DeleteById_InternalServerError() {
	var (
		id, _       = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
		url         = fmt.Sprintf("%s/%s", s.baseURI, id.String())
		request, _  = http.NewRequest(http.MethodDelete, url, nil)
		recorder    = httptest.NewRecorder()
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), name, description))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "pizzas" WHERE "pizzas"."id" = $1`)).
		WithArgs(id).
		WillReturnError(errors.New("something didn't work"))
	s.mock.ExpectRollback()

	s.router.ServeHTTP(recorder, request)
	s.T().Run("should return status code 500", func(t *testing.T) {
		assert.Equal(s.T(), http.StatusInternalServerError, recorder.Code)
	})
}
