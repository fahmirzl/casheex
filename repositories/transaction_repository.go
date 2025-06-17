package repositories

import (
	"casheex/structs"
	"database/sql"
)

func InsertTransaction(dbParam *sql.DB, transaction *structs.Transaction) error {
	sqlStatement := "INSERT INTO transactions VALUES(null, DATE(NOW()), ?, ?, ?, ?, NOW(), null)"
	_, err := dbParam.Exec(sqlStatement, transaction.UserID, transaction.Total, transaction.Paid,transaction.Change)
	if err != nil {
		return err
	}

	err = InsertTransactionDetail(dbParam, transaction)
	if err != nil {
		return err
	}

	var user structs.User
	user.ID = *transaction.UserID

	err = DeductStockOnTransaction(dbParam, &user)
	if err != nil {
		return err
	}

	err = TruncateCart(dbParam, &user)

	return err
}