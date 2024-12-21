package broker

import "encoding/json"

type NewPostMessage struct {
	AuthorId string `json:"author_id"`
}

type Message interface {
	ToValue() (value []byte, err error)
}

func (m *NewPostMessage) ToValue() ([]byte, error) {
	value, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return value, nil
}
