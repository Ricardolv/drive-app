package queue

type QueueResponse struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Id       int    `json:"id"`
}
