package repository

import (
	"fmt"
	"log"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ShoppingCartRepository struct {
	DBConn *sqlx.DB
}

func NewShoppingCartRepository(DBConn *sqlx.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{
		DBConn: DBConn,
	}
}

func (scr *ShoppingCartRepository) Create(userId int, cart model.ShoppingCart) error {

	return nil
}

func (scr *ShoppingCartRepository) Browse(userId int) (model.ShoppingCart, error) {
	var (
		cart model.ShoppingCart
	)

	cartStatement := `
		SELECT id, user_id, total
		FROM shopping_carts
		WHERE
			user_id = $1
	`

	err := scr.DBConn.QueryRowx(cartStatement, userId).StructScan(&cart)
	if err != nil {
		log.Print(fmt.Errorf("error ShoppingCartRepository - Browse: %w", err))
		return cart, err
	}

	itemStatement := `
		SELECT id, shopping_cart_id, product_id, quantity
		FROM shopping_cart_items
		WHERE
			shopping_cart_id = $1
	`

	rows, err := scr.DBConn.Queryx(itemStatement, cart.ID)
	if err != nil {
		log.Print(fmt.Errorf("error ShoppingCartRepository - Browse: %w", err))
		return cart, err
	}

	for rows.Next() {
		var item model.ShoppingCartItem
		rows.StructScan(&item)
		cart.Items = append(cart.Items, item)
	}

	return cart, nil
}

func (scr *ShoppingCartRepository) GetItemById(userId int, itemId int) (model.ShoppingCartItem, error) {
	var (
		item      model.ShoppingCartItem
		statement = `
			SELECT id, shopping_cart_id, product_id, quantity
			FROM shopping_cart_items
			LEFT JOIN shopping_carts
				ON shopping_carts.id = shopping_cart_id
			WHERE
				shopping_carts.user_id = $1
				shopping_cart_items.id = $2
		`
	)

	err := scr.DBConn.QueryRowx(statement, userId, itemId).StructScan(&item)
	if err != nil {
		log.Print(fmt.Errorf("error ShoppingCartRepository - GetItemById: %w", err))
		return item, err
	}

	return item, nil
}

func (scr *ShoppingCartRepository) Update(userId int, item model.ShoppingCartItem) error {
	return nil
}

func (scr *ShoppingCartRepository) Delete(userId int, itemId int) error {
	return nil
}
