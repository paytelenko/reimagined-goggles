package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint) (Task, error)
	GetTasksByUserID(userID uint) ([]Task, error)
	UpdateTask(task Task) error
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}
func (r *taskRepository) CreateTask(task *Task) (*Task, error) {
	err := r.db.Create(task).Error
	return task, err
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uint) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
