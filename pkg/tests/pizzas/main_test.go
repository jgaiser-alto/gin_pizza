package pizzatests

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pizza/pkg/pizzas"
	"testing"
)

type PizzaTestSuite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	router     *gin.Engine
	repository pizzas.Repository
	baseURI    string
}

func (s *PizzaTestSuite) SetupTest() {
	// prevent shared state between tests
	s.SetupSuite()
}

func (s *PizzaTestSuite) AfterTest(_, _ string) {
	defer require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestPizzaTestSuite(t *testing.T) {
	suite.Run(t, new(PizzaTestSuite))
}

func (s *PizzaTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = pizzas.CreateRepository(s.DB)
	s.baseURI = "/pizzas"

	router := gin.Default()
	pizzas.RegisterRoutes(router, s.DB)
	s.router = router
}
