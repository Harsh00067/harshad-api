package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Harsh00067/harshad-api/models"
	"github.com/Harsh00067/harshad-api/repository"
)

func CreateUser(user models.User) error {
	fmt.Println("Service: CreateUser is called")
	fmt.Println("User in service:", user)

	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.DOJ = strings.TrimSpace(user.DOJ)

	if user.Name == "" {
		return errors.New("name is required")
	}

	if len(user.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email format")
	}

	if user.DOJ == "" {
		return errors.New("DOJ is required")
	}

	doj, err := time.Parse("2006-01-02", user.DOJ)
	if err != nil {
		return errors.New("invalid date format, use YYYY-MM-DD")
	}

	if doj.After(time.Now()) {
		return errors.New("DOJ cannot be in the future")
	}

	if doj.Year() < 2000 {
		return errors.New("DOJ is too old")
	}

	return repository.CreateUser(user)
}

func SearchUsers(name string, page int, limit int) ([]models.User, error) {

	fmt.Println("Service: SearchUsers is called")
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit
	if name == "" {
		return repository.GetUsers(limit, offset)
	}

	return repository.SearchUsers(name, limit, offset)
}

func GetUserByID(id int) (models.User, error) {
	fmt.Println("Service: GetUsersById is called")
	return repository.GetUserByID(id)
}
func GetUser() ([]models.User, error) {
	fmt.Println("Service: GetUsers is called")
	return repository.FetchUsers()
}

func DeleteUser(id int) error {
	fmt.Println("Service: DeleteUser is called")
	if id <= 0 {
		return errors.New("Id must be greater than or equal to zero")
	}
	return repository.DeleteUser(id)
}
func UpdateUser(id int, user models.User) error {
	fmt.Println("Service:UpdateUser is called")
	if id <= 0 {
		return errors.New("Id must be greater than or equal to zero")
	}
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.DOJ = strings.TrimSpace(user.DOJ)

	if user.Name == "" {
		return errors.New("name is required")
	}

	if len(user.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email format")
	}

	return repository.UpdateUser(id, user)
}
