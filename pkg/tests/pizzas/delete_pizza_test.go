package tests

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"pizza/pkg/common/models"
	"regexp"
)

func (s *PizzaTestSuite) TestExpectedPizzaIsDeleted() {
	var (
		id, _       = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
	)
	url := fmt.Sprintf("%s/%s", s.baseUri, id.String())
	request, _ := http.NewRequest(http.MethodDelete, url, nil)
	recorder := httptest.NewRecorder()

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

	// Convert the JSON response to a map
	var response models.Pizza
	json.Unmarshal([]byte(recorder.Body.String()), &response)

	assert.Equal(s.T(), http.StatusOK, recorder.Code)
}
