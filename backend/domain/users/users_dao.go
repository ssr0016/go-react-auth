package users

import (
	"react-auth/backend/datasource/mysql/users_db"
	"react-auth/backend/utils/errors"
)

// DAO data access object
// DAO will be used to interact with the database

var (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, password) VALUES(?, ?, ?, ?)"
)

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewBadRequestError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save user")
	}
	user.ID = userID

	return nil
}
