package queue

import "encoding/json"

type Message struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	ID       int    `json:"id"`
}

func (response *Message) Marshal() ([]byte, error) {
	return json.Marshal(response)
}

func (response *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, response)
}
