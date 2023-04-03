package repository

import (
	"fmt"
	"log"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserAddressRepository struct {
	DBConn *sqlx.DB
}

func NewUserAddressRepository(DBConn *sqlx.DB) *UserAddressRepository {
	return &UserAddressRepository{
		DBConn: DBConn,
	}
}

func (cr *UserAddressRepository) Create(userAddress model.UserAddress) error {
	statement := `
		INSERT INTO user_addresses(user_id, address_line, country, city, postal_code, phone, note, is_main)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id

		`
	lastId := 0

	err := cr.DBConn.QueryRow(statement,
		userAddress.UserId, userAddress.AddressLine, userAddress.Country, userAddress.City,
		userAddress.PostalCode, userAddress.Phone, userAddress.Note, userAddress.IsMain,
	).Scan(&lastId)

	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - Create: %w", err))
		return err
	}

	if userAddress.IsMain {
		err = cr.UpdateMainAddress(lastId, userAddress.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *UserAddressRepository) Browse(userId int) ([]model.UserAddress, error) {
	var (
		userAddresses []model.UserAddress
		statement     = `
			SELECT id, user_id, address_line, country, city, postal_code, phone, note, is_main
			FROM user_addresses
			WHERE
				user_id = $1
		`
	)

	rows, err := cr.DBConn.Queryx(statement, userId)
	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - Browse : %w", err))
		return userAddresses, err
	}

	for rows.Next() {
		var user model.UserAddress
		rows.StructScan(&user)
		userAddresses = append(userAddresses, user)
	}

	return userAddresses, nil
}

func (cr *UserAddressRepository) GetById(id int, userId int) (model.UserAddress, error) {
	var (
		userAddress model.UserAddress
		statement   = `
			SELECT id, user_id, address_line, country, city, postal_code, phone, note, is_main
			FROM user_addresses
			WHERE
				id = $1 AND
				user_id = $2
			LIMIT 1
		`
	)

	err := cr.DBConn.QueryRowx(statement, id, userId).StructScan(&userAddress)
	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - GetById: %w", err))
		return userAddress, err
	}

	return userAddress, nil
}

func (cr *UserAddressRepository) Update(userAddress model.UserAddress) error {
	statement := `
		UPDATE user_addresses
		SET
			address_line = $2, country = $3, city = $4, 
			postal_code = $5, phone = $6, note = $7, is_main = $8, updated_at = CURRENT_TIMESTAMP
		WHERE
			user_id = $1 AND
			id = $9
	`
	_, err := cr.DBConn.Queryx(statement,
		userAddress.UserId, userAddress.AddressLine, userAddress.Country, userAddress.City,
		userAddress.PostalCode, userAddress.Phone, userAddress.Note, userAddress.IsMain,
		userAddress.ID,
	)

	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - Update: %w", err))
		return err
	}

	if userAddress.IsMain {
		err = cr.UpdateMainAddress(userAddress.ID, userAddress.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *UserAddressRepository) Delete(id int, userId int) error {
	statement := `
		UPDATE user_addresses
		SET
			updated_at = CURRENT_TIMESTAMP, deleted_at = CURRENT_TIMESTAMP
		WHERE 
			id = $1 AND 
			user_id $2
	`
	_, err := cr.DBConn.Queryx(statement, id, userId)
	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - Delete: %w", err))
		return err
	}

	return nil
}

func (cr *UserAddressRepository) UpdateMainAddress(id int, userId int) error {
	statement := `
		UPDATE user_addresses
		SET 
			is_main = false, updated_at = CURRENT_TIMESTAMP
		WHERE
			id != $1 AND
			user_id = $2
	`
	log.Print("Current ID: ", id)
	_, err := cr.DBConn.Queryx(statement, id, userId)
	if err != nil {
		log.Print(fmt.Errorf("error UserAddressRepository - UpdateMainAddress: %w", err))
		return err
	}

	return nil
}
