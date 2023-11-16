package mockclient

type MockClient struct {
	Data string
}

func (m *MockClient) Send() ([]byte, error) {
	data := m.Data

	if len(data) == 0 {
		data = "send"
	}

	return []byte(data), nil
}
