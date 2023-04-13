package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/models"

	"github.com/jmoiron/sqlx"
)

type PostgresDBDiscountRepository struct {
	db *sqlx.DB
}

func NewPostgresDBDiscountRepository(db *sqlx.DB) DiscountRepository {

	return &PostgresDBDiscountRepository{db: db}
}

func (r *PostgresDBDiscountRepository) GetDiscount(ID int) (*models.Discount, error) {
	discount := &models.Discount{}

	err := r.db.Get(discount, "SELECT * FROM Discount WHERE ID = $1", ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get discount: %w", err)
	}

	return discount, nil
}

func (r *PostgresDBDiscountRepository) CreateDiscount(discount *models.Discount) (*models.Discount, error) {

	result, err := r.db.NamedExec(
		`INSERT INTO Discount (product_id, product_name, description, discount_type, percentage, amount, quantity, start_date, end_date, created_at, updated_at) 
		VALUES (:product_id, :product_name, :description, :discount_type, :percentage, :amount, :quantity, :start_date, :end_date, :created_at, :updated_at)
		RETURN id
		`, discount)

	if err != nil {
		return nil, fmt.Errorf("failed to create discount: %w", err)
	}

	discountID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to create discount: %w", err)
	}

	create_discount := &models.Discount{}

	err = r.db.Get(create_discount, "SELECT * FROM Discount WHERE ID = $1", discountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get discount: %w", err)
	}

	return create_discount, nil
}

func (r *PostgresDBDiscountRepository) UpdateDiscount(discount *models.Discount) (*models.Discount, error) {

	result, err := r.db.NamedExec(`UPDATE Discount SET
	product_id = :product_id, 
	product_name = :product_name, 
	description = :description, 
	percentage = :percentage, 
	amount = :amount, 
	quantity = :quantity, 
	start_date = :start_date, 
	end_date = :end_date, 
	updated_at = :updated_at
	WHERE id= :id
	RETURNING id
	`, discount)

	if err != nil {
		return nil, fmt.Errorf("failed to update discount: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	update_discount := &models.Discount{}

	err = r.db.Get(update_discount, "SELECT * FROM Discount WHERE ID = $1", discount.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get discount: %w", err)
	}

	return update_discount, nil
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
