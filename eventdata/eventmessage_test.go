package eventdata

import "testing"

func TestEventMessageIsReady(t *testing.T) {
	msg := EventMessage{}

	if msg.IsReady() {
		t.Error("event message should not be ready")
	}

	msg = EventMessage{
		Message: "Testing",
		Event:   "testevent.test",
		Type:    "test",
		Data:    map[string]interface{}{"example": "value"},
	}

	if !msg.IsReady() {
		t.Error("event message should be ready")
	}
}

func TestEventMessageToJson(t *testing.T) {
	msg := EventMessage{
		Message: "Testing",
		Event:   "testevent.test",
		Type:    "test",
		Data:    map[string]interface{}{"example": "value"},
	}

	m, err := msg.ToJson()

	if err != nil {
		t.Error(err)
	}

	expected := `{"message":"Testing","event":"testevent.test","context":"test","data":{"example":"value"}}`

	if string(m) != expected {
		t.Errorf("Expected %s, not equals to %s", expected, string(m))
	}
}

func TestEventMessageToMap(t *testing.T) {
	msg := EventMessage{
		Message: "Testing",
		Event:   "testevent.test",
		Type:    "test",
		Data:    map[string]interface{}{"example": "value"},
	}

	m, err := msg.ToMap()

	if err != nil {
		t.Error(err)
	}

	if m["message"] != "Testing" || m["event"] != "testevent.test" || m["context"] != "test" {
		t.Errorf("Event message values might be incorrect, %v", m)
	}
}
