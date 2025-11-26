package userService

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	User string `json:"user"`
}
