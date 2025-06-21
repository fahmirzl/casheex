package repositories

import (
	"casheex/structs"
	"database/sql"
)

func InsertTransaction(dbParam *sql.DB, transaction *structs.Transaction) error {
	sqlStatement := "INSERT INTO transactions VALUES(null, DATE(NOW()), ?, ?, ?, ?, NOW(), null)"
	_, err := dbParam.Exec(sqlStatement, transaction.UserID, transaction.Total, transaction.Paid, transaction.Change)
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

func GetTransactionByUserIdAndDate(dbParam *sql.DB, user *structs.User, date string) (result []structs.TransactionResponse, err error) {
	var sqlStatement string
	var transactions *sql.Rows

	if date == "" {
		sqlStatement = "SELECT * FROM transactions WHERE user_id = ?"
		transactions, err = dbParam.Query(sqlStatement, user.ID)
	} else {
		sqlStatement = "SELECT * FROM transactions WHERE user_id = ? AND date = ?"
		transactions, err = dbParam.Query(sqlStatement, user.ID, date)
	}
	if err != nil {
		panic(err)
	}
	defer transactions.Close()

	for transactions.Next() {
		var transactionResponse structs.TransactionResponse
		err = transactions.Scan(&transactionResponse.ID, &transactionResponse.Date, &transactionResponse.UserID, &transactionResponse.Total, &transactionResponse.Paid, &transactionResponse.Change, &transactionResponse.CreatedAt, &transactionResponse.UpdatedAt)
		if err != nil {
			panic(err)
		}
		var transactionDetailsResponse []structs.TransactionDetailResponse
		sqlSelectStatement := "SELECT id, product_id, purchase_price, selling_price, quantity, subtotal, created_at, updated_at FROM transaction_details WHERE transaction_id = ?"
		transactionDetails, err := dbParam.Query(sqlSelectStatement, transactionResponse.ID)
		if err != nil {
			panic(err)
		}
		defer transactionDetails.Close()

		for transactionDetails.Next() {
			var transactionDetail structs.TransactionDetailResponse
			err = transactionDetails.Scan(&transactionDetail.ID, &transactionDetail.ProductID, &transactionDetail.PurchasePrice, &transactionDetail.SellingPrice, &transactionDetail.Quantity, &transactionDetail.Subtotal, &transactionDetail.CreatedAt, &transactionDetail.UpdatedAt)
			if err != nil {
				panic(err)
			}

			var user structs.User
			var userResponse structs.UserResponse
			user.ID = *transactionResponse.UserID
			err = GetUserById(dbParam, &user, &userResponse)
			if err != nil {
				panic(err)
			}
			transactionResponse.User = userResponse

			var product structs.Product
			product.ID = *transactionDetail.ProductID
			err = GetProductById(dbParam, &product)
			if err != nil {
				panic(err)
			}

			transactionDetail.Product = product

			transactionDetailsResponse = append(transactionDetailsResponse, transactionDetail)
		}
		transactionResponse.TransactionDetails = transactionDetailsResponse
		result = append(result, transactionResponse)
	}

	return
}

func GetAllTransaction(dbParam *sql.DB, date string) (result []structs.TransactionResponse, err error) {
	var sqlStatement string
	var transactions *sql.Rows

	if date == "" {
		sqlStatement = "SELECT * FROM transactions"
		transactions, err = dbParam.Query(sqlStatement)
	} else {
		sqlStatement = "SELECT * FROM transactions WHERE date = ?"
		transactions, err = dbParam.Query(sqlStatement, date)
	}
	if err != nil {
		panic(err)
	}
	defer transactions.Close()

	for transactions.Next() {
		var transactionResponse structs.TransactionResponse
		err = transactions.Scan(&transactionResponse.ID, &transactionResponse.Date, &transactionResponse.UserID, &transactionResponse.Total, &transactionResponse.Paid, &transactionResponse.Change, &transactionResponse.CreatedAt, &transactionResponse.UpdatedAt)
		if err != nil {
			panic(err)
		}

		var user structs.User
		var userResponse structs.UserResponse
		user.ID = *transactionResponse.UserID
		err = GetUserById(dbParam, &user, &userResponse)
		if err != nil {
			panic(err)
		}
		transactionResponse.User = userResponse

		var transactionDetailsResponse []structs.TransactionDetailResponse
		sqlSelectStatement := "SELECT id, product_id, purchase_price, selling_price, quantity, subtotal, created_at, updated_at FROM transaction_details WHERE transaction_id = ?"
		transactionDetails, err := dbParam.Query(sqlSelectStatement, transactionResponse.ID)
		if err != nil {
			panic(err)
		}
		defer transactionDetails.Close()

		for transactionDetails.Next() {
			var transactionDetail structs.TransactionDetailResponse
			err = transactionDetails.Scan(&transactionDetail.ID, &transactionDetail.ProductID, &transactionDetail.PurchasePrice, &transactionDetail.SellingPrice, &transactionDetail.Quantity, &transactionDetail.Subtotal, &transactionDetail.CreatedAt, &transactionDetail.UpdatedAt)
			if err != nil {
				panic(err)
			}

			var product structs.Product
			product.ID = *transactionDetail.ProductID
			err = GetProductById(dbParam, &product)
			if err != nil {
				panic(err)
			}

			transactionDetail.Product = product

			transactionDetailsResponse = append(transactionDetailsResponse, transactionDetail)
		}
		transactionResponse.TransactionDetails = transactionDetailsResponse
		result = append(result, transactionResponse)
	}

	return
}

func GetProfitBetweenDate(dbParam *sql.DB, profitResponse *structs.ProfitResponse, startDate string, endDate string) error {
	sqlStatement := "SELECT SUM(transaction_details.selling_price * transaction_details.quantity) AS revenue, SUM(transaction_details.purchase_price * transaction_details.quantity) AS expense, SUM(transaction_details.selling_price * transaction_details.quantity) - SUM(transaction_details.purchase_price * transaction_details.quantity) AS profit, 'IDR' as currency FROM transactions INNER JOIN transaction_details ON transactions.id = transaction_details.transaction_id WHERE transactions.date BETWEEN ? AND ?"
	err := dbParam.QueryRow(sqlStatement, startDate, endDate).Scan(&profitResponse.Revenue, &profitResponse.Expense, &profitResponse.Profit, &profitResponse.Currency)
	return err
}