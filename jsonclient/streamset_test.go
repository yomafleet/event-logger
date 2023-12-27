package jsonclient

import (
	"testing"
)

func TestStreamAddLable(t *testing.T) {
	s := StreamSet{}
	expect := "value"
	label := map[string]string{"key": expect}

	s.AddLabel(&label)

	if s.Stream["key"] != "value" {
		t.Fatalf("Label value should be %s, instead of %s", s.Stream["key"], expect)
	}
}

func TestStreamAddValue(t *testing.T) {
	s := StreamSet{}
	value := map[string]interface{}{"SetA": "A", "SetB": "B"}
	stream, err := s.AddValue(&value)

	if err != nil || len(stream.Values) != 1 {
		t.Fatal("Value not inserted")
	}

	t.Logf("Values: %v", s.Values)
}

func TestStreamToJson(t *testing.T) {
	s := StreamSet{
		Stream: map[string]string{"key": "value"},
		Values: [][]string{{"1234", "a"}, {"1235", "b"}},
	}

	marshalled, err := s.ToJson()

	if err != nil {
		t.Error(err)
	}

	json := string(marshalled)
	want := `{"stream":{"key":"value"},"values":[["1234","a"],["1235","b"]]}`

	if want != json {
		t.Fatalf("Expected: %s, not equals to %s", want, json)
	}
}
