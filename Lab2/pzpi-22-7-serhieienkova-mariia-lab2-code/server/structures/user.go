package structures

type User struct {
	Id           int    `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	Name         string `json:"name" db:"name"`
	Surname      string `json:"surname" db:"surname"`
	PasswordHash string `json:"password" db:"password_hash"`
}

type UserToken struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

type UserType string
