package jsonclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	netUrl "net/url"
)

type httpClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type JsonClient struct {
	Url        string
	Streams    map[string]StreamSet
	HttpClient httpClient
}

func (j *JsonClient) WithHttpClient(url string) (*JsonClient, error) {
	j.HttpClient = &http.Client{}

	return j, nil
}

func (j *JsonClient) AddStream(key string, stream *StreamSet) *JsonClient {
	existed, ok := j.Streams[key]

	if ok {
		existed.Values = append(existed.Values, stream.Values...)
		j.Streams[key] = existed

		return j
	}

	if j.Streams == nil {
		j.Streams = map[string]StreamSet{}
	}

	j.Streams[key] = *stream

	return j
}

func (j *JsonClient) GetStream(key string) (*StreamSet, error) {
	stream, ok := j.Streams[key]

	if !ok {
		return nil, errors.New("stream is not set yet")
	}

	return &stream, nil
}

func (j *JsonClient) AppendValue(key string, val *map[string]interface{}) (*JsonClient, error) {
	stream, streamErr := j.GetStream(key)

	if streamErr != nil {
		return nil, streamErr
	}

	appended, err := stream.AddValue(val)

	if err != nil {
		return nil, err
	}

	j.AddStream(key, appended)

	return j, err
}

func (j *JsonClient) ToJson() ([]byte, error) {
	var streamList []StreamSet

	for _, stream := range j.Streams {
		streamList = append(streamList, stream)
	}

	streams := map[string][]StreamSet{
		"streams": streamList,
	}

	marshalled, err := json.Marshal(streams)

	if err != nil {
		return nil, err
	}

	return marshalled, nil
}

func (j *JsonClient) Send() ([]byte, error) {
	_, err := netUrl.ParseRequestURI(j.Url)

	if err != nil {
		return nil, err
	}

	if j.HttpClient == nil {
		j.WithHttpClient(j.Url)
	}

	jsonBytes, err := j.ToJson()

	if err != nil {
		return nil, err
	}

	response, err := j.HttpClient.Post(j.Url, "application/json", bytes.NewBuffer(jsonBytes))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	j.Flush()

	return body, err
}

func (j *JsonClient) Flush() {
	j.Streams = nil
}
