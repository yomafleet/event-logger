package eventdata

import (
	"encoding/json"
)

type EventMessage struct {
	Message string                 `json:"message"`
	Event   string                 `json:"event"`
	Type    string                 `json:"type"`
	Data    map[string]interface{} `json:"data"`
}

func (e *EventMessage) IsReady() bool {
	return len(e.Message) > 0 && len(e.Event) > 0 && len(e.Data) > 0
}

func (s *EventMessage) ToJson() ([]byte, error) {
	msg, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *EventMessage) ToMap() (map[string]interface{}, error) {
	marshalled, err := s.ToJson()

	if err != nil {
		return nil, err
	}

	var mapped map[string]interface{}
	err = json.Unmarshal(marshalled, &mapped)

	if err != nil {
		return nil, err
	}

	return mapped, nil
}
