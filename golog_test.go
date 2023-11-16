package elog

import (
	"testing"

	"github.com/yomafleet/elog/eventdata"
)

func TestNewLoggerSend(t *testing.T) {
	logger := New("")
	logger = logger.NewWithClient("mock")

	err := logger.Send(eventdata.EventMessage{
		Message: "Testing",
		Event:   "testing.event",
		Context: "testing",
		Data:    map[string]interface{}{"testing": true},
	})

	if err != nil {
		t.Error(err)
	}
}
