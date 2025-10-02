package domain

type User struct {
	ID          int    `json:"id" db:"id"`
	UserName    string `json:"user_name" db:"user_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password,omitempty" db:"password"`
	IsShowOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}
