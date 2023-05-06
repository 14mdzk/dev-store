package repository

import (
	"fmt"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DBConn *sqlx.DB
}

func NewUserRepository(DBConn *sqlx.DB) *UserRepository {
	return &UserRepository{
		DBConn: DBConn,
	}
}

func (cr *UserRepository) Create(user model.User) error {
	statement := `
		INSERT INTO users(username, email, password)
		VALUES ($1, $2, $3)
	`
	_, err := cr.DBConn.Queryx(statement, user.Username, user.Email, user.Password)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create: %w", err))
		return err
	}

	return nil
}

func (cr *UserRepository) Browse() ([]model.User, error) {
	var (
		users     []model.User
		statement = `
			SELECT id, username, email
			FROM users
		`
	)

	rows, err := cr.DBConn.Queryx(statement)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Browse : %w", err))
		return users, err
	}

	for rows.Next() {
		var user model.User
		rows.StructScan(&user)
		users = append(users, user)
	}

	return users, nil
}

func (cr *UserRepository) GetById(id int) (model.User, error) {
	var (
		user      model.User
		statement = `
			SELECT id, username, email
			FROM users
			WHERE
				id = $1
			LIMIT 1
		`
	)

	err := cr.DBConn.QueryRowx(statement, id).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetById: %w", err))
		return user, err
	}

	return user, nil
}

func (cr *UserRepository) GetByEmail(email string) (model.User, error) {
	var (
		user      model.User
		statement = `
			SELECT id, username, email, password
			FROM users
			WHERE
				email = $1
			LIMIT 1
		`
	)

	err := cr.DBConn.QueryRowx(statement, email).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmail: %w", err))
		return user, err
	}

	return user, nil
}

func (cr *UserRepository) Update(user model.User) error {
	statement := `
		UPDATE users
		SET
			username = $1, email = $2, password = $3, updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $4
	`
	_, err := cr.DBConn.Queryx(statement, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Update: %w", err))
		return err
	}

	return nil
}

func (cr *UserRepository) Delete(id int) error {
	statement := `
		UPDATE users
		SET
			updated_at = CURRENT_TIMESTAMP, deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err := cr.DBConn.Queryx(statement, id)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Delete: %w", err))
		return err
	}

	return nil
}
