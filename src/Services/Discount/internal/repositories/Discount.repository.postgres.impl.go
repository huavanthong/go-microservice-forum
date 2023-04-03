package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresDBDiscountRepository struct {
	db *sqlx.DB
}

func NewPostgresDBDiscountRepository(cfg config.DatabaseConfig) (*PostgresDBDiscountRepository, error) {

	// Create connection string from config
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	// Open connection on PostgreSQL
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}
	// Ping database to ensure connection is valid
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &PostgresDBDiscountRepository{
		db: db,
	}, nil
}

func (r *PostgresDBDiscountRepository) GetDiscount(productName string) (*models.Coupon, error) {
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

func (r *PostgresDBDiscountRepository) CreateDiscount(coupon *models.Coupon) error {
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

func (r *PostgresDBDiscountRepository) UpdateDiscount(coupon *models.Coupon) error {
	result, err := r.db.Exec(
		"UPDATE Coupon SET ProductName=$1, Description=$2, Amount=$3 WHERE Id=$4",
		coupon.ProductName, coupon.Description, coupon.Amount, coupon.ID,
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

func (r *PostgresDBDiscountRepository) DeleteDiscount(productName string) error {
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
