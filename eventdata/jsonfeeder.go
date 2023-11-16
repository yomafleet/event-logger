package eventdata

import (
	"errors"

	"github.com/yomafleet/elog/jsonclient"
)

type JsonFeeder struct {
	message EventMessage
	client  *jsonclient.JsonClient
}

func (j *JsonFeeder) SetMessage(msg *EventMessage) {
	j.message = *msg
}

func (j *JsonFeeder) SetClient(client *jsonclient.JsonClient) *JsonFeeder {
	j.client = client

	return j
}

func (j *JsonFeeder) Feed() error {
	if j.client == nil {
		return errors.New("client has not been set")
	}

	stream, err := j.mapToStreamSet()

	if err != nil {
		return err
	}

	key := j.message.Context + "_" + j.message.Event
	j.client.AddStream(key, stream)

	return nil
}

func (j *JsonFeeder) mapToStreamSet() (*jsonclient.StreamSet, error) {
	if !j.message.IsReady() {
		return nil, errors.New("event message is not ready, it might be empty")
	}

	label := map[string]string{"event": j.message.Context}
	stream := jsonclient.StreamSet{}
	stream.AddLabel(&label)

	mapped, err := j.message.ToMap()

	if err != nil {
		return nil, err
	}

	stream.AddValue(&mapped)

	return &stream, nil
}
