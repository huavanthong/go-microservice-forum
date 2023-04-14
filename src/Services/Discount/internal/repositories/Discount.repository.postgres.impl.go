package repositories

import (
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"

	"gorm.io/gorm"
)

type PostgresDBDiscountRepository struct {
	db *gorm.DB
}

func NewPostgresDBDiscountRepository(db *gorm.DB) DiscountRepository {

	return &PostgresDBDiscountRepository{db: db}
}

func (r *PostgresDBDiscountRepository) GetDiscount(id string) (*models.Discount, error) {
	discount := &models.Discount{}

	if err := r.db.First(discount, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Discount not found with id: %s", id)
		}
		// Internal error on db
		return nil, fmt.Errorf("Error on DB find for user", err)
	}

	return discount, nil
}

func (r *PostgresDBDiscountRepository) CreateDiscount(discount *models.Discount) (*models.Discount, error) {

	if err := r.db.Create(discount).Error; err != nil {
		return nil, fmt.Errorf("Faield to create discount", err)
	}

	create_discount := &models.Discount{ID: discount.ID}

	if err := r.db.First(create_discount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Discount not found with id", discount.ID)
		}
		// Internal error on db
		return nil, fmt.Errorf("Error on DB find for user", err)
	}

	return create_discount, nil
}

func (r *PostgresDBDiscountRepository) UpdateDiscount(discount *models.Discount) (*models.Discount, error) {

	temDiscount := &models.Discount{ID: discount.ID}

	if err := r.db.First(temDiscount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Discount not found with id", discount.ID)
		}
		// Internal error on db
		return nil, fmt.Errorf("Error on DB find for user", err)
	}

	if err := r.db.Save(discount).Error; err != nil {
		return nil, fmt.Errorf("Error saving user: ", err)
	}

	return discount, nil
}

func (r *PostgresDBDiscountRepository) DeleteDiscount(id string) error {

	discount := &models.Discount{}
	if err := r.db.First(discount, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Discount not found with id", discount.ID)
		}
		// Internal error on db
		return fmt.Errorf("Error on DB find for user", err)
	}

	if err := r.db.Delete(discount).Error; err != nil {
		return fmt.Errorf("Error deleting user", err)
	}
	return nil
}
