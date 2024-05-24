package queue

import "encoding/json"

type QueueResponse struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	ID       int    `json:"id"`
}

func (response *QueueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(response)
}

func (response *QueueResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, response)
}
