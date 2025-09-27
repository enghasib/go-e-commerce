package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	UserName    string `json:"user_name" db:"user_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password,omitempty" db:"password"`
	IsShowOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(User User) (*User, error)
	Get(userId int) (*User, error)
	List() ([]*User, error)
	Update(id int, User User) (*User, error)
	Delete(userId int) error
	Find(email, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {

	query := `
		INSERT INTO users(user_name, email, password, is_shop_owner) VALUES(
		$1, $2, $3, $4
		)
		RETURNING id
	`

	row := r.db.QueryRow(query, user.UserName, user.Email, user.Password, user.IsShowOwner)

	if row.Err() != nil {
		fmt.Println("err", row.Err())
		return nil, row.Err()
	}

	row.Scan(&user.ID)

	return &user, nil
}

func (r *userRepo) Get(userId int) (*User, error) {
	// for _, user := range r.userList {
	// 	if user.ID == userId {
	// 		return user, nil
	// 	}
	// }
	// return nil, errors.New("user not found!")
	var user *User

	query := `
		SELECT * FROM users WHERE id=$1
	`
	row := r.db.QueryRow(query, userId)
	row.Scan(&user)
	return user, nil

}

func (r *userRepo) List() ([]*User, error) {
	var userList []*User

	query := `SELECT id, user_name, email, is_shop_owner FROM users`
	err := r.db.Select(&userList, query)
	if err != nil {
		fmt.Println("err:", err)
	}

	return userList, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	var user User
	query := `
		SELECT id, user_name, email, password, is_shop_owner 
		FROM users 
		WHERE email=$1 AND password=$2

	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	return &user, nil

}

func (r *userRepo) Update(id int, user User) (*User, error) {
	// for i := range r.userList {
	// 	if r.userList[i].ID == id {
	// 		if &user.UserName != nil {
	// 			r.userList[i].UserName = user.UserName
	// 		}
	// 		if &user.Email != nil {
	// 			r.userList[i].Email = user.Email
	// 		}
	// 	}
	// 	return r.userList[i], nil
	// }
	return nil, errors.New("update failed!")
}

func (r *userRepo) Delete(userId int) error {
	// index := -1

	// for i, product := range r.userList {
	// 	if product.ID == userId {
	// 		index = i
	// 		break
	// 	}
	// }
	// r.userList = append(r.userList[:index], r.userList[index+1:]...)
	return nil
}
