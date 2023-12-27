package eventdata

type MockFeeder struct {
	Data error
}

func (m *MockFeeder) SetMessage(msg *EventMessage) {
	//
}

func (m *MockFeeder) Feed() error {
	if m.Data != nil {
		return m.Data
	}

	return nil
}
