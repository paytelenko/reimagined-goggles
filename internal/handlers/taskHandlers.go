package handlers

import (
	"awesomeProject/internal/taskService"
	"awesomeProject/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	Service taskService.TaskService
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}
func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{Service: s}
}

func (h *TaskHandler) PatchTasks(ctx context.Context, request tasks.PatchTasksRequestObject) (tasks.PatchTasksResponseObject, error) {
	id := *request.Body.Id
	task := *request.Body
	updatedTask, err := h.Service.UpdateTask(id, taskService.Task{
		Text:   *task.Task,
		IsDone: *task.IsDone,
	})
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasks200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id
	task, err := h.Service.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Text,
		IsDone: &task.IsDone,
	}, nil
}
