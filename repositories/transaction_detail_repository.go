package repositories

import (
	"casheex/structs"
	"database/sql"
)

func InsertTransactionDetail(dbParam *sql.DB, transaction *structs.Transaction) error {
	sqlStatement := `
	INSERT INTO transaction_details
	SELECT
		null,
		(SELECT id FROM transactions WHERE user_id = ? ORDER BY id DESC LIMIT 1),
		carts.product_id,
		products.purchase_price,
		carts.selling_price,
		carts.quantity,
		carts.subtotal,
		NOW(),
		null
	FROM carts
	INNER JOIN products ON carts.product_id = products.id
	WHERE carts.user_id = ?;
	`
	_, err := dbParam.Exec(sqlStatement, transaction.UserID, transaction.UserID)
	return err
}