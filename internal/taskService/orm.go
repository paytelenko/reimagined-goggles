package taskService

type Task struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Text   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
type TaskRequest struct {
	Task string `json:"task"`
}
