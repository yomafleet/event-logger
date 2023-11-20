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

type JsonClientConfig struct {
	Url   string
	Id    string
	Token string
}

type JsonClient struct {
	Config     JsonClientConfig
	Streams    map[string]*StreamSet
	HttpClient httpClient
}

func (j *JsonClient) WithHttpClient() (*JsonClient, error) {
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
		j.Streams = map[string]*StreamSet{}
	}

	j.Streams[key] = stream

	return j
}

func (j *JsonClient) GetStream(key string) (*StreamSet, error) {
	stream, ok := j.Streams[key]

	if !ok {
		return nil, errors.New("stream is not set yet")
	}

	return stream, nil
}

func (j *JsonClient) AppendValue(key string, val *map[string]interface{}) (*JsonClient, error) {
	stream, streamErr := j.GetStream(key)

	if streamErr != nil {
		return nil, streamErr
	}

	_, err := stream.AddValue(val)

	if err != nil {
		return nil, err
	}

	return j, err
}

func (j *JsonClient) ToJson() ([]byte, error) {
	var streamList []StreamSet

	for _, stream := range j.Streams {
		streamList = append(streamList, *stream)
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
	url, err := j.buildUrlFromConfig()

	if err != nil {
		return nil, err
	}

	if j.HttpClient == nil {
		j.WithHttpClient()
	}

	jsonBytes, err := j.ToJson()

	if err != nil {
		return nil, err
	}

	response, err := j.HttpClient.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	j.Clear()

	return body, err
}

func (j *JsonClient) Clear() {
	j.Streams = nil
}

func (j *JsonClient) buildUrlFromConfig() (string, error) {
	parsed, err := netUrl.ParseRequestURI(j.Config.Url)

	if err != nil {
		return "", err
	}

	if len(j.Config.Id) > 0 && len(j.Config.Token) > 0 {
		parsed.User = netUrl.UserPassword(j.Config.Id, j.Config.Token)
	}

	return parsed.String(), nil
}
