package jsonclient

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type StreamSet struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}

func (s *StreamSet) AddLabel(label *map[string]string) *StreamSet {
	s.Stream = *label

	return s
}

func (s *StreamSet) AddValue(val *map[string]interface{}) (*StreamSet, error) {
	value, err := makeValue(val)

	if err != nil {
		return nil, err
	}

	s.Values = append(s.Values, value)

	return s, nil
}

func (s *StreamSet) ToJson() ([]byte, error) {
	stream, err := json.Marshal(s)

	if err != nil {
		log.Printf("Stream set cannot be cannot marshal: %s \n ", err)

		return nil, err
	}

	return stream, nil
}

func makeValue(line *map[string]interface{}) ([]string, error) {
	logline, err := json.Marshal(*line)

	if err != nil {
		log.Printf("Logline cannot marshal: %s \n ", err)

		return nil, err
	}

	timestamp := time.Now().UnixNano()

	return []string{fmt.Sprint(timestamp), string(logline)}, nil
}
