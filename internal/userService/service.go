package userService

import "awesomeProject/internal/taskService"

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]*User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(id uint, user *User) (*User, error)
	DeleteUser(id uint) error
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}
type userService struct {
	repo        UserRepository
	taskService taskService.TaskService
}

func NewUserService(r UserRepository, ts taskService.TaskService) UserService {
	return &userService{repo: r, taskService: ts}
}

func (s *userService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return s.taskService.GetTasksByUserID(userID)
}

func (s *userService) CreateUser(user *User) (*User, error) {
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]*User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id uint, user *User) (*User, error) {
	u, err := s.repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	u.Email = user.Email
	u.Password = user.Password

	if err := s.repo.UpdateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
