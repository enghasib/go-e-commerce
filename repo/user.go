package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/enghasib/server/domain"
	"github.com/enghasib/server/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {

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

func (r *userRepo) Get(userId int) (*domain.User, error) {
	// for _, user := range r.userList {
	// 	if user.ID == userId {
	// 		return user, nil
	// 	}
	// }
	// return nil, errors.New("user not found!")
	var user *domain.User

	query := `
		SELECT * FROM users WHERE id=$1
	`
	row := r.db.QueryRow(query, userId)
	row.Scan(&user)
	return user, nil

}

func (r *userRepo) List(limit, page int) ([]*domain.User, error) {

	offset := ((page - 1) * limit) + 1

	var userList []*domain.User

	query := `SELECT id, user_name, email, is_shop_owner FROM users LIMIT $1 OFFSET $2`
	err := r.db.Select(&userList, query, limit, offset)
	if err != nil {
		fmt.Println("err:", err)
	}

	return userList, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT id, user_name, email, password, is_shop_owner 
		FROM users 
		WHERE email=$1 AND password=$2

	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		fmt.Println(err)
		return nil, err
	}

	return &user, nil

}

func (r *userRepo) Update(id int, user domain.User) (*domain.User, error) {
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
	return nil, errors.New("update failed")
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
