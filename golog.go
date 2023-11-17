package elog

import (
	"errors"

	"github.com/yomafleet/elog/eventdata"
	"github.com/yomafleet/elog/jsonclient"
	"github.com/yomafleet/elog/mockclient"
)

type logClient interface {
	Send() ([]byte, error)
}

type logFeeder interface {
	SetMessage(msg *eventdata.EventMessage)
	Feed() error
}

type Logger struct {
	config *Config
	client *logClient
	feeder *logFeeder
}

func (l *Logger) AddMessage(msg *eventdata.EventMessage) error {
	feeder := *l.feeder
	feeder.SetMessage(msg)
	err := feeder.Feed()

	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) Send(msg *eventdata.EventMessage) error {
	if msg != nil {
		err := l.AddMessage(msg)

		if err != nil {
			return err
		}
	}

	client := *l.client
	client.Send()

	return nil
}

func (l *Logger) NewWithClient(client string) *Logger {
	return makeLogger(&Config{
		Client:   client,
		Settings: l.config.Settings,
	})
}

func New(configPath string) *Logger {
	return makeLogger(loadConfig(configPath))
}

func makeLogger(config *Config) *Logger {
	client, feeder, err := setup(config)

	if err != nil {
		panic(err)
	}

	return &Logger{client: &client, feeder: &feeder, config: config}
}

func loadConfig(path string) *Config {
	if len(path) == 0 {
		return MustLoadConfig("./config.yaml")
	}

	c, err := LoadConfig(path)

	if err != nil {
		panic(err)
	}

	return c
}

func setup(config *Config) (logClient, logFeeder, error) {
	if config.Client == "json" {
		client := jsonclient.JsonClient{Url: config.Settings["json"]["url"]}
		feeder := eventdata.JsonFeeder{}
		feeder.SetClient(&client)

		return &client, &feeder, nil
	} else if config.Client == "mock" {
		return &mockclient.MockClient{}, &eventdata.MockFeeder{}, nil
	}

	return nil, nil, errors.New("no client nor feeder found")
}
