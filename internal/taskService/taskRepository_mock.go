package taskService

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository - поддельный репозиторий
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *Task) (*Task, error) {
	args := m.Called(task)
	var t *Task
	if res := args.Get(0); res != nil {
		t = res.(*Task)
	}
	return t, args.Error(1)
}

func (m *MockTaskRepository) GetAllTasks() ([]Task, error) {
	args := m.Called()
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id uint) (Task, error) {
	args := m.Called(id)
	var t Task
	if res := args.Get(0); res != nil {
		t = res.(Task)
	}
	return t, args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
