package service

import (
	"connect_mysql/database"
	"connect_mysql/model"
	"log"
)

// UserService định nghĩa các phương thức mà một service xử lý User cần phải triển khai.
type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) GetUserByID(id int) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	var u model.User
	err = row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	insertStmt, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	updateStmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer updateStmt.Close()
	_, err = updateStmt.Exec(user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil

}

func (s *UserServiceImpl) DeleteUser(id int) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	deleteStmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer deleteStmt.Close()
	_, err = deleteStmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
