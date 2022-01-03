package structs

type Stats struct {
	Free   string `json:"free"`
	Df     string `json:"df"`
	Uptime string `json:"uptime"`
	Ps     string `json:"ps"`
}
