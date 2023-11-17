package elog

import (
	"testing"

	"github.com/yomafleet/eventlogger/eventdata"
)

func TestNewLoggerAddMessage(t *testing.T) {
	logger := New("")
	logger = logger.NewWithClient("mock")

	err := logger.AddMessage(&eventdata.EventMessage{
		Message: "Testing",
		Event:   "testing.event",
		Context: "testing",
		Data:    map[string]interface{}{"testing": true},
	})

	if err != nil {
		t.Error(err)
	}

	err = logger.Send(nil)

	if err != nil {
		t.Error(err)
	}
}

func TestNewLoggerSend(t *testing.T) {
	logger := New("")
	logger = logger.NewWithClient("mock")

	err := logger.Send(&eventdata.EventMessage{
		Message: "Testing",
		Event:   "testing.event",
		Context: "testing",
		Data:    map[string]interface{}{"testing": true},
	})

	if err != nil {
		t.Error(err)
	}
}
