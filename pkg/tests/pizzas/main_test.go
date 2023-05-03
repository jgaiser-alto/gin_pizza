package tests

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pizza/pkg/common/models"
	"pizza/pkg/pizzas"
	"testing"
)

type PizzaTestSuite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	router     *gin.Engine
	repository pizzas.Repository
	pizza      *models.Pizza
	baseUri    string
}

func (s *PizzaTestSuite) BeforeTest(_, _ string) {
	// prevent shared state between tests
	s.SetupSuite()
}

func (s *PizzaTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
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
	s.baseUri = "/pizzas"

	router := gin.Default()
	pizzas.RegisterRoutes(router, s.DB)
	s.router = router
}
