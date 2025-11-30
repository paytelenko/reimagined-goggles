package taskService

type Task struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Text   string `json:"task"`
	IsDone bool   `json:"isDone"`
	UserID uint   `json:"user_id"`
}
type TaskRequest struct {
	Task string `json:"task"`
}
