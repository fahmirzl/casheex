package repositories

import (
	"casheex/structs"
	"database/sql"
)

func GetAllProduct(dbParam *sql.DB) (result []structs.Product, err error) {
	sqlStatement := "SELECT * FROM products"
	rows, err := dbParam.Query(sqlStatement)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product structs.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Stock, &product.PurchasePrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return
		}
		result = append(result, product)
	}
	return
}

func InsertProduct(dbParam *sql.DB, product *structs.Product) error {
	sqlStatement := "INSERT INTO products VALUES(null, ?, ?, ?, ?, NOW(), null)"
	result, err := dbParam.Exec(sqlStatement, product.Name, product.Stock, product.PurchasePrice, product.SellingPrice)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	sqlSelectStatement := "SELECT * FROM products WHERE id = ?"
	err = dbParam.QueryRow(sqlSelectStatement, id).Scan(&product.ID, &product.Name, &product.Stock, &product.PurchasePrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
	return err
}

func GetProductById(dbParam *sql.DB, product *structs.Product) error {
	sqlStatement := "SELECT * FROM products WHERE id = ?"
	err := dbParam.QueryRow(sqlStatement, product.ID).Scan(&product.ID, &product.Name, &product.Stock, &product.PurchasePrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
	return err
}

func UpdateProduct(dbParam *sql.DB, product *structs.Product) error {
	sqlStatement := "UPDATE products SET name = ?, stock = ?, purchase_price = ?, selling_price = ?, updated_at = NOW() WHERE id = ?"
	result, err := dbParam.Exec(sqlStatement, product.Name, product.Stock, product.PurchasePrice, product.SellingPrice, product.ID)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	sqlSelectStatement := "SELECT * FROM products WHERE id = ?"
	err = dbParam.QueryRow(sqlSelectStatement, product.ID).Scan(&product.ID, &product.Name, &product.Stock, &product.PurchasePrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
	return err
}

func DeleteProduct(dbParam *sql.DB, product *structs.Product) error {
	sqlStatement := "DELETE FROM products WHERE id = ?"
	result, err := dbParam.Exec(sqlStatement, product.ID)
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	return err
}