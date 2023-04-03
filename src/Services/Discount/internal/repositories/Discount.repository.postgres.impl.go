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

func (r *PostgresDBDiscountRepository) GetDiscount(ID int) (*models.DiscountDBResponse, error) {
	discount := &models.DiscountDBResponse{}

	err := r.db.Get(discount, "SELECT * FROM Discount WHERE ID = $1", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get discount: %w", err)
	}

	return discount, nil
}

func (r *PostgresDBDiscountRepository) CreateDiscount(discount *models.Discount) error {
	result, err := r.db.Exec(
		"INSERT INTO Discount (ProductName, Description, Percentage, Quantity, StartDate, EndDate) VALUES ($1, $2, $3, $4, $5, $6)",
		discount.ProductName, discount.Description, discount.Percentage, discount.Quantity, discount.StartDate, discount.EndDate,
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

func (r *PostgresDBDiscountRepository) UpdateDiscount(discount *models.Discount) error {
	result, err := r.db.Exec(
		"UPDATE Coupon SET ProductName=$1, Description=$2, Percentage=$3, Quantity=$4, StartDate=$5, EndDate=$6 WHERE ID=$7",
		discount.ProductName, discount.Description, discount.Percentage, discount.Quantity, discount.StartDate, discount.EndDate, discount.ID,
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

func (r *PostgresDBDiscountRepository) DeleteDiscount(ID int) error {
	result, err := r.db.Exec("DELETE FROM Discount WHERE ID = $1", ID)
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
