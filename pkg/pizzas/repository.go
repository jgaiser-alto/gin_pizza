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
	Delete(pizza models.Pizza) error
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
	var pizzas []*models.Pizza
	result := p.DB.Find(&pizzas)
	if result.Error != nil {
		fmt.Printf("pizza machine broke: %d", result.Error)
		return nil, result.Error
	}

	return pizzas, result.Error
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
	result := p.DB.Clauses(clause.Returning{}).Select("Name", "Description").Save(&pizza)
	if result.Error != nil {
		fmt.Printf("pizza machine broke: %d", result.Error)
		return nil, result.Error
	}
	return &pizza, result.Error
}

func (p *repo) Delete(pizza models.Pizza) error {
	result := p.DB.Delete(&pizza)
	if result.Error != nil {
		fmt.Printf("pizza machine broke: %d", result.Error)
		return result.Error
	}
	return result.Error
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{DB: db}
}
