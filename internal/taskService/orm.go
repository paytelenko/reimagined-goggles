package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone string `json:"isDone"`
}
type TaskRequest struct {
	Task string `json:"task"`
}
