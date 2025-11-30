package userService

import "awesomeProject/internal/taskService"

type User struct {
	ID       uint               `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	User string `json:"user"`
}
