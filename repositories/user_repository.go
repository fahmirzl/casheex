package repositories

import (
	"casheex/configs"
	"casheex/structs"
	"database/sql"
)

func CheckLogin(dbParam *sql.DB, user *structs.User) error {
	sqlStatement := "SELECT id, username, role FROM users WHERE username=? AND password=?"
	err := dbParam.QueryRow(sqlStatement, user.Username, configs.MD5(user.Password)).Scan(&user.ID, &user.Username, &user.Role)
	return err
}

func GetAllUser(dbParam *sql.DB) (result []structs.UserResponse, err error) {
	sqlStatement := "SELECT id, name, gender, username, role, created_at, updated_at FROM users"
	rows, err := dbParam.Query(sqlStatement)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.UserResponse
		err = rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return
		}
		result = append(result, user)
	}
	return
}

func InsertUser(dbParam *sql.DB, user *structs.User, userResponse *structs.UserResponse) error {
	sqlStatement := "INSERT INTO users VALUES(null, ?, ?, ?, ?, ?, NOW(), null)"
	result, err := dbParam.Exec(sqlStatement, user.Name, user.Gender, user.Username, configs.MD5(user.Password), user.Role)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	sqlSelectStatement := "SELECT id, name, gender, username, role, created_at, updated_at FROM users WHERE id = ?"
	err = dbParam.QueryRow(sqlSelectStatement, id).Scan(&userResponse.ID, &userResponse.Name, &userResponse.Gender, &userResponse.Username, &userResponse.Role, &userResponse.CreatedAt, &userResponse.UpdatedAt)
	return err
}

func GetUserById(dbParam *sql.DB, user *structs.User, userResponse *structs.UserResponse) error {
	sqlStatement := "SELECT id, name, gender, username, role, created_at, updated_at FROM users WHERE id = ?"
	err := dbParam.QueryRow(sqlStatement, user.ID).Scan(&userResponse.ID, &userResponse.Name, &userResponse.Gender, &userResponse.Username, &userResponse.Role, &userResponse.CreatedAt, &userResponse.UpdatedAt)
	return err
}

func UpdateUser(dbParam *sql.DB, user *structs.User, userResponse *structs.UserResponse) error {
	sqlStatement := "UPDATE users SET name = ?, gender = ?, username = ?, password = ?, role = ?, updated_at = NOW() WHERE id = ?"
	result, err := dbParam.Exec(sqlStatement, user.Name, user.Gender, user.Username, configs.MD5(user.Password), user.Role, user.ID)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	sqlSelectStatement := "SELECT id, name, gender, username, role, created_at, updated_at FROM users WHERE id = ?"
	err = dbParam.QueryRow(sqlSelectStatement, user.ID).Scan(&userResponse.ID, &userResponse.Name, &userResponse.Gender, &userResponse.Username, &userResponse.Role, &userResponse.CreatedAt, &userResponse.UpdatedAt)
	return err
}

func DeleteUser(dbParam *sql.DB, user *structs.User) error {
	sqlStatement := "DELETE FROM users WHERE id = ?"
	result, err := dbParam.Exec(sqlStatement, user.ID)
	row, _ := result.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	return err
}