package structs

type Task struct {
	TaskID  string `json:"task_id"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}
