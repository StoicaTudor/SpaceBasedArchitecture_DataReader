package repository

import (
	"DataReader/data"
	"DataReader/database_connector"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetUser(id string) (data.User, error) {
	// START DB initialization
	logErrorMessage := "Error getting user: %v"
	appErrorMessage := "cannot get user"

	db, dbInitError := database_connector.InitDB()
	if dbInitError != nil {
		log.Printf(logErrorMessage, dbInitError)
		return data.User{}, errors.New(appErrorMessage)
	}
	defer database_connector.CloseDB()
	// END DB initialization

	query := "SELECT * FROM user WHERE id = ?"
	row := db.QueryRow(query, id)

	var user data.User
	err := row.Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return data.User{}, fmt.Errorf("user with ID %s not found", id)
		}
		log.Printf("Error retrieving user: %v", err)
		return data.User{}, err
	}
	return user, nil
}

func GetAllUsers() ([]data.User, error) {
	// START DB initialization
	logErrorMessage := "Error getting all user: %v"
	appErrorMessage := "cannot get all users"

	db, dbInitError := database_connector.InitDB()
	if dbInitError != nil {
		log.Printf(logErrorMessage, dbInitError)
		return nil, errors.New(appErrorMessage)
	}
	defer database_connector.CloseDB()
	// END DB initialization

	// START query
	var users []data.User
	query := "SELECT * FROM user"
	rows, dbQueryError := db.Query(query)
	if dbQueryError != nil {
		log.Printf("Error executing query: %v", dbQueryError)
		return users, dbQueryError
	}
	defer rows.Close()
	// END query

	for rows.Next() {
		var user data.User
		err := rows.Scan(&user.ID, &user.Name, &user.Balance)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return users, nil
}
