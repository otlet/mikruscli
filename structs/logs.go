package structs

type Logs struct {
	ID          string `json:"id"`
	ServerID    string `json:"server_id"`
	Task        string `json:"task"`
	WhenCreated string `json:"when_created"`
	WhenDone    string `json:"when_done"`
	Output      string `json:"output"`
}
