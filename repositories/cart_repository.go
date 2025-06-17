package repositories

import (
	"casheex/structs"
	"database/sql"
)

func CartWithProductByUserId(dbParam *sql.DB, cart *structs.Cart) (result []structs.CartResponse, err error) {
	sqlStatement := "SELECT id, product_id, selling_price, quantity, subtotal, created_at, updated_at FROM carts WHERE user_id = ?"
	rows, err := dbParam.Query(sqlStatement, cart.UserID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cartResponse structs.CartResponse
		var product structs.Product

		err = rows.Scan(&cartResponse.ID, &cartResponse.ProductID, &cartResponse.SellingPrice, &cartResponse.Quantity, &cartResponse.Subtotal, &cartResponse.CreatedAt, &cartResponse.UpdatedAt)
		if err != nil {
			return
		}
		sqlSelectStatement := "SELECT * FROM products WHERE id = ?"
		err = dbParam.QueryRow(sqlSelectStatement, cartResponse.ProductID).Scan(&product.ID, &product.Name, &product.Stock, &product.PurchasePrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
		cartResponse.Product = product
		result = append(result, cartResponse)
	}
	return
}

func AddProductToCart(dbParam *sql.DB, cart *structs.Cart) error {
	var tempCart structs.Cart
	sqlCheckStatement := "SELECT * FROM carts WHERE product_id = ? AND user_id = ?"
	err := dbParam.QueryRow(sqlCheckStatement, cart.ProductID, cart.UserID).Scan(
		&tempCart.ID, &tempCart.ProductID, &tempCart.SellingPrice, &tempCart.Quantity,
		&tempCart.Subtotal, &tempCart.UserID, &tempCart.CreatedAt, &tempCart.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		sqlInsertStatement := "INSERT INTO carts VALUES(null, ?, ?, ?, ?, ?, NOW(), null)"
		_, err := dbParam.Exec(sqlInsertStatement,
			cart.ProductID, cart.SellingPrice, cart.Quantity, cart.Subtotal, cart.UserID,
		)
		return err
	} else if err != nil {
		return err
	} else {
		newQuantity := *tempCart.Quantity + *cart.Quantity
		newSubtotal := newQuantity * *cart.SellingPrice
		sqlUpdateStatement := "UPDATE carts SET quantity = ?, subtotal = ?, updated_at = NOW() WHERE id = ?"
		_, err := dbParam.Exec(sqlUpdateStatement, newQuantity, newSubtotal, tempCart.ID)
		return err
	}
}

func DeleteProductFromCart(dbParam *sql.DB, cart *structs.Cart) error {
	sqlStatement := "DELETE FROM carts WHERE id = ?"
	result, err := dbParam.Exec(sqlStatement, cart.ID)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	return err
}

func TruncateCart(dbParam *sql.DB, user *structs.User) error {
	sqlStatement := "DELETE FROM carts WHERE user_id = ?"
	_, err := dbParam.Exec(sqlStatement, user.ID)
	return err
}