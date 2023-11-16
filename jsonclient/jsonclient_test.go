package jsonclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonClientAddStream(t *testing.T) {
	j := JsonClient{}
	streamset := StreamSet{
		Stream: map[string]string{"label": "value"},
		Values: [][]string{{"12345", "abcdef"}, {"12346", "abcdeg"}},
	}

	j.AddStream("example", &streamset)

	_, ok := j.Streams["example"]

	if !ok {
		t.Error("Stream list is empty")
	}
}

func TestJsonClientGetStream(t *testing.T) {
	j := JsonClient{}
	streamset := StreamSet{
		Stream: map[string]string{"label": "value"},
		Values: [][]string{{"12345", "abcdef"}, {"12346", "abcdeg"}},
	}

	j.AddStream("example", &streamset)

	_, err := j.GetStream("example")

	if err != nil {
		t.Error(err)
	}
}

func TestJsonClientAppendValue(t *testing.T) {
	j := JsonClient{}
	streamset := StreamSet{
		Stream: map[string]string{"label": "value"},
		Values: [][]string{{"12345", "abcdef"}, {"12346", "abcdeg"}},
	}

	j.AddStream("example", &streamset)

	newValue := map[string]interface{}{"12347": "abcdeh"}

	j.AppendValue("example", &newValue)

	stream, _ := j.GetStream("example")

	if len(stream.Values) != 3 {
		t.Errorf("streams length count is: %d instead of expected, %d", len(stream.Values), 3)
	}
}

func TestJsonClientToJson(t *testing.T) {
	j := JsonClient{}
	streamset := StreamSet{
		Stream: map[string]string{"label": "value"},
		Values: [][]string{{"12345", "abcdef"}, {"12346", "abcdeg"}},
	}

	j.AddStream("example", &streamset)

	marshalled, err := j.ToJson()

	if err != nil {
		t.Error(err)
	}

	json := string(marshalled)
	want := `{"streams":[{"stream":{"label":"value"},"values":[["12345","abcdef"],["12346","abcdeg"]]}]}`

	if want != json {
		t.Fatalf("Expected: %s, not equals to %s", want, json)
	}
}

func TestJsonClientSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))

	defer server.Close()

	j := JsonClient{Url: server.URL}
	streamset := StreamSet{
		Stream: map[string]string{"label": "value"},
		Values: [][]string{{"12345", "abcdef"}, {"12346", "abcdeg"}},
	}

	j.AddStream("example", &streamset)

	body, err := j.Send()

	if err != nil {
		t.Error(err)
	}

	if string(body) != `{"status":"success"}` {
		t.Errorf("Expected: %s, not equals to %s", `{"status":"success"}`, string(body))
	}
}
