package repository

import (
	"fmt"
	"log"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CategoryRepository struct {
	DBConn *sqlx.DB
}

func NewCategoryRepository(DBConn *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		DBConn: DBConn,
	}
}

func (cr *CategoryRepository) Create(category model.Category) error {
	statement := `
		INSERT INTO categories(name, description)
		VALUES ($1, $2)
	`
	_, err := cr.DBConn.Queryx(statement, category.Name, category.Description)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Create: %w", err))
		return err
	}

	return nil
}

func (cr *CategoryRepository) Browse() ([]model.Category, error) {
	var (
		categories []model.Category
		statement  = `
			SELECT id, name, description
			FROM categories
		`
	)

	rows, err := cr.DBConn.Queryx(statement)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		rows.StructScan(&category)
		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *CategoryRepository) GetById(id int) (model.Category, error) {
	var (
		category  model.Category
		statement = `
			SELECT id, name, description
			FROM categories
			WHERE
				id = $1
			LIMIT 1
		`
	)

	err := cr.DBConn.QueryRowx(statement, id).StructScan(&category)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - GetById: %w", err))
		return category, err
	}

	return category, nil
}

func (cr *CategoryRepository) Update(category model.Category) error {
	statement := `
		UPDATE categories
		SET
			name = $1, description = $2
		WHERE
			id = $3
	`
	_, err := cr.DBConn.Queryx(statement, category.Name, category.Description, category.ID)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Update: %w", err))
		return err
	}

	return nil
}

func (cr *CategoryRepository) Delete(id int) error {
	statement := `
		DELETE FROM categories
		WHERE id = $1
	`
	_, err := cr.DBConn.Queryx(statement, id)
	if err != nil {
		log.Print(fmt.Errorf("error CategoryRepository - Delete: %w", err))
		return err
	}

	return nil
}
