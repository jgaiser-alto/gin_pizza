package pizza_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"pizza/pkg/common/models"
	"regexp"
)

func (s *PizzaTestSuite) TestExpectedPizzaIsReturned() {
	var (
		id, _       = uuid.NewUUID()
		name        = "test-name"
		description = "a test pizza"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "pizzas" WHERE (id = $1)`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id.String(), name, description))

	res, err := s.repo.Get(id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&models.Pizza{ID: id, Name: name}, res))
}
