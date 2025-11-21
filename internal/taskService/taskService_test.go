package taskService

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     *Task
		mockSetup func(m *MockTaskRepository, input *Task)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: &Task{Text: "Test", IsDone: false},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(input, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании",
			input: &Task{Text: "Bad task", IsDone: false},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(&Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewTasksService(mockRepo)
			result, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	tests := []struct {
		name         string
		initialValue []Task
		mockSetup    func(m *MockTaskRepository, initialValue []Task)
		wantErr      bool
	}{
		{

			name:         "Получение пустого списка",
			initialValue: []Task{},
			mockSetup: func(m *MockTaskRepository, initialValue []Task) {
				m.On("GetAllTasks").Return(initialValue, nil)
			},
			wantErr: false,
		},
		{
			name:         "Получение заполненного списка",
			initialValue: []Task{{Text: "Test_1", IsDone: false}, {Text: "Test_2", IsDone: true}},
			mockSetup: func(m *MockTaskRepository, initialValue []Task) {
				m.On("GetAllTasks").Return(initialValue, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.initialValue)

			service := NewTasksService(mockRepo)
			result, err := service.GetAllTasks()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.initialValue, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}

}

func TestUpdateTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		input     Task
		mockSetup func(m *MockTaskRepository, id uint, input Task)
		want      Task
		wantErr   bool
	}{
		{name: "Успешное обновление задачи",
			id:    1,
			input: Task{ID: 1, Text: "Test_1", IsDone: true},
			mockSetup: func(m *MockTaskRepository, id uint, input Task) {
				updatedTask := Task{ID: id, Text: input.Text, IsDone: input.IsDone}
				m.On("GetTaskByID", id).Return(updatedTask, nil)
				m.On("UpdateTask", input).Return(nil)
			},
			want:    Task{ID: 1, Text: "Test_1", IsDone: true},
			wantErr: false,
		},
		{name: "Ошибка при обновлении задачи",
			id:    2,
			input: Task{ID: 2, Text: "Test_2", IsDone: false},
			mockSetup: func(m *MockTaskRepository, id uint, input Task) {
				m.On("GetTaskByID", id).Return(Task{
					ID:     2,
					Text:   "",
					IsDone: false,
				}, nil)
				m.On("UpdateTask", input).Return(errors.New("invalid task ID"))
			},
			want:    Task{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id, tt.input)

			service := NewTasksService(mockRepo)
			result, err := service.UpdateTask(tt.id, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		mockSetup func(m *MockTaskRepository, id uint)
		wantErr   bool
	}{
		{name: "Успешное удаление задачи",
			id: 1,
			mockSetup: func(m *MockTaskRepository, id uint) {
				m.On("DeleteTask", id).Return(nil)
			},
			wantErr: false,
		},
		{name: "Ошибка при удаление задачи",
			id: 2,
			mockSetup: func(m *MockTaskRepository, id uint) {
				m.On("DeleteTask", id).Return(errors.New("invalid task ID"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewTasksService(mockRepo)
			err := service.DeleteTask(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
