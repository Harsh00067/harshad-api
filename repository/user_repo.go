package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Harsh00067/harshad-api/database"
	"github.com/Harsh00067/harshad-api/models"
)

func CreateUser(user models.User) error {
	fmt.Println("Repository: CreateUser is called")
	now := time.Now()
	_, err := database.DB.Exec(
		"INSERT INTO users(name,email,doj,created_at,updated_at) VALUES($1,$2,$3,$4,$5)",
		user.Name,
		user.Email,
		user.DOJ,
		now,
		now,
	)

	return err
}

func SearchUsers(name string, limit int, offset int) ([]models.User, error) {
	fmt.Println("Repository: SearchUsers is called")
	rows, err := database.DB.Query(
		"SELECT id,name,email,doj FROM users WHERE name ILIKE '%' || $1 || '%' ORDER BY id LIMIT $2 OFFSET $3", name, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.DOJ,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("no user found")
	}

	return users, nil
}

func GetUsers(limit int, offset int) ([]models.User, error) {
	fmt.Println("Repository: SearchUsers is called")
	rows, err := database.DB.Query(
		"SELECT id,name,email,doj,created_at,updated_at FROM users ORDER BY id LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.DOJ,
			&user.Created_AT,
			&user.Updated_AT,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("no user found")
	}

	return users, nil
}

func GetUserByID(id int) (models.User, error) {

	fmt.Println("Repository: GetUserByID is called")

	var user models.User

	err := database.DB.QueryRow(
		"SELECT id,name,email,doj FROM users WHERE id=$1",
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.DOJ,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func DeleteUser(id int) error {
	fmt.Println("Repository: DeleteUser is called")
	result, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)

	if err != nil {
		return err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsaffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func UpdateUser(id int, user models.User) error {
	fmt.Println("Repository:UpdateUser is called")
	now := time.Now()
	result, err := database.DB.Exec("UPDATE users SET name=$1, email=$2, doj=$3, updated_at=$4 WHERE id=$5", user.Name, user.Email, user.DOJ, now, id)
	if err != nil {
		return err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsaffected == 0 {
		return errors.New("user not found")
	}
	return nil

}

func FetchUsers() ([]models.User, error) {
	fmt.Println("Repository: FetchUsers is called")
	rows, err := database.DB.Query(
		"SELECT id,name,email,doj,created_at,updated_at FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.DOJ,
			&user.Created_AT,
			&user.Updated_AT,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.New("no user found")
	}

	return users, nil
}
