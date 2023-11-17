package eventdata

import (
	"testing"

	"github.com/yomafleet/eventlogger/jsonclient"
)

func TestJsonFeederFeed(t *testing.T) {
	msg := EventMessage{
		Message: "Testing",
		Event:   "testevent.test",
		Context: "test",
		Data:    map[string]interface{}{"example": "value"},
	}

	j := JsonFeeder{}
	j.SetMessage(&msg)
	j.SetClient(&jsonclient.JsonClient{})

	err := j.Feed()

	if err != nil {
		t.Error(err)
	}
}
