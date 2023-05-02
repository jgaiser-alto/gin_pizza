package pizzas

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"pizza/pkg/common/models"
)

type Repository interface {
	Get(id uuid.UUID) (*models.Pizza, error)
	GetAll() ([]*models.Pizza, error)
	Create(pizza models.Pizza) (*models.Pizza, error)
	Update(pizza models.Pizza) (*models.Pizza, error)
	Delete(id uuid.UUID) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) Get(id uuid.UUID) (*models.Pizza, error) {
	var pizza models.Pizza
	result := p.DB.First(&pizza, id)
	if result.Error != nil {
		fmt.Printf("pizza machine broke: %d", result.Error)
		return nil, result.Error
	}
	return &pizza, result.Error
}

func (p *repo) GetAll() ([]*models.Pizza, error) {
	//TODO implement me
	panic("implement me")
}

func (p *repo) Create(pizza models.Pizza) (*models.Pizza, error) {
	result := p.DB.Clauses(clause.Returning{}).Select("Name", "Description").Create(&pizza)
	if result.Error != nil {
		fmt.Printf("pizza machine broke: %d", result.Error)
		return nil, result.Error
	}
	return &pizza, result.Error
}

func (p *repo) Update(pizza models.Pizza) (*models.Pizza, error) {
	//TODO implement me
	panic("implement me")
}

func (p *repo) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{DB: db}
}
