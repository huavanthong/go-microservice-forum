package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type DiscountRepository struct {
	db *sqlx.DB
}

func NewDiscountRepository(config *viper.Viper) (*DiscountRepository, error) {
	// Get connection string to PostgreSQL
	connString := config.GetString("DatabaseSettings.ConnectionString")
	// Open connection on PostgreSQL
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &DiscountRepository{
		db: db,
	}, nil
}

func (r *DiscountRepository) GetDiscount(productName string) (*models.Coupon, error) {
	coupon := &models.Coupon{}

	err := r.db.Get(coupon, "SELECT * FROM Coupon WHERE ProductName = $1", productName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get discount: %w", err)
	}

	return coupon, nil
}

func (r *DiscountRepository) CreateDiscount(coupon *models.Coupon) error {
	result, err := r.db.Exec(
		"INSERT INTO Coupon (ProductName, Description, Amount) VALUES ($1, $2, $3)",
		coupon.ProductName, coupon.Description, coupon.Amount,
	)
	if err != nil {
		return fmt.Errorf("failed to create discount: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *DiscountRepository) UpdateDiscount(coupon *models.Coupon) error {
	result, err := r.db.Exec(
		"UPDATE Coupon SET ProductName=$1, Description=$2, Amount=$3 WHERE Id=$4",
		coupon.ProductName, coupon.Description, coupon.Amount, coupon.Id,
	)
	if err != nil {
		return fmt.Errorf("failed to update discount: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (r *DiscountRepository) DeleteDiscount(productName string) error {
	result, err := r.db.Exec("DELETE FROM Coupon WHERE ProductName = $1", productName)
	if err != nil {
		return fmt.Errorf("failed to delete discount: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
