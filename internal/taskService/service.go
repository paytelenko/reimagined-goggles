package taskService

type TaskService interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskByID(id uint) (Task, error)
	UpdateTask(id uint, task Task) (Task, error)
	DeleteTask(id uint) error
	GetTasksByUserID(userID uint) ([]Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTasksService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task *Task) (*Task, error) {
	return s.repo.CreateTask(task)
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *taskService) GetTaskByID(id uint) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id uint, task Task) (Task, error) {
	tsk, err := s.repo.GetTaskByID(id)

	if err != nil {
		return Task{}, err
	}

	tsk.Text = task.Text
	tsk.IsDone = task.IsDone

	if err := s.repo.UpdateTask(tsk); err != nil {
		return Task{}, err
	}
	return tsk, nil
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
