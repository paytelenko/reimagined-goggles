package taskService

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(task string) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id, task string) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTasksService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task string) (Task, error) {
	tsk := Task{
		ID:     uuid.NewString(),
		Task:   task,
		IsDone: "in progress",
	}
	if err := s.repo.CreateTask(tsk); err != nil {
		return Task{}, err
	}
	return tsk, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id, task string) (Task, error) {
	tsk, err := s.repo.GetTaskByID(id)

	if err != nil {
		return Task{}, err
	}

	tsk.Task = task
	tsk.IsDone = "done"

	if err := s.repo.UpdateTask(tsk); err != nil {
		return Task{}, err
	}
	return tsk, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
