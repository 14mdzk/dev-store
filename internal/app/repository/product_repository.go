package repository

import (
	"fmt"
	"log"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ProductRepository struct {
	DBConn *sqlx.DB
}

func NewProductRepository(DBConn *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		DBConn: DBConn,
	}
}

func (cr *ProductRepository) Create(Product model.Product) error {
	statement := `
		INSERT INTO products(category_id, name, description, price, stock, is_active, currency)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := cr.DBConn.Queryx(statement, Product.CategoryId, Product.Name, Product.Description, fmt.Sprintf("%0.2f", Product.Price), Product.Stock, Product.IsActive, Product.Currency)
	if err != nil {
		log.Print(fmt.Errorf("error ProductRepository - Create: %w", err))
		return err
	}

	return nil
}

func (cr *ProductRepository) Browse() ([]model.Product, error) {
	var (
		products  []model.Product
		statement = `
			SELECT id, category_id, name, description, currency, price, stock, is_active
			FROM products
		`
	)

	rows, err := cr.DBConn.Queryx(statement)
	if err != nil {
		log.Print(fmt.Errorf("error ProductRepository - Browse : %w", err))
		return products, err
	}

	for rows.Next() {
		var product model.Product
		rows.StructScan(&product)
		products = append(products, product)
	}

	return products, nil
}

func (cr *ProductRepository) GetById(id int) (model.Product, error) {
	var (
		product   model.Product
		statement = `
			SELECT id, category_id, name, description, currency, price, stock, is_active
			FROM products
			WHERE
				id = $1
			LIMIT 1
		`
	)

	err := cr.DBConn.QueryRowx(statement, id).StructScan(&product)
	if err != nil {
		log.Print(fmt.Errorf("error ProductRepository - GetById: %w", err))
		return product, err
	}

	return product, nil
}

func (cr *ProductRepository) Update(product model.Product) error {
	statement := `
		UPDATE products
		SET
			category_id = $1, name = $2, description = $3, price = $4, stock = $5, is_active = $6,
			currency = $7, updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $8
	`

	log.Print(product)
	_, err := cr.DBConn.Queryx(statement, product.CategoryId, product.Name, product.Description, fmt.Sprintf("%0.2f", product.Price), product.Stock, product.IsActive, product.Currency, product.ID)
	if err != nil {
		log.Print(fmt.Errorf("error ProductRepository - Update: %w", err))
		return err
	}

	return nil
}

func (cr *ProductRepository) Delete(id int) error {
	statement := `
		UPDATE products
		SET
			updated_at = CURRENT_TIMESTAMP, deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err := cr.DBConn.Queryx(statement, id)
	if err != nil {
		log.Print(fmt.Errorf("error ProductRepository - Delete: %w", err))
		return err
	}

	return nil
}
