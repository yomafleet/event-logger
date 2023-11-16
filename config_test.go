package elog

import (
	"testing"
)

func TestConfigMustLoadConfig(t *testing.T) {
	c := MustLoadConfig("./config.yaml")

	if len(c.Client) == 0 {
		t.Error("config not set")
	}

	if c.Client == "json" && len(c.Settings["json"]["url"]) == 0 {
		t.Error("Config for 'json' client has no 'url' field")
	}
}

func TestConfigLoadConfig(t *testing.T) {
	c, err := LoadConfig("./config.yaml")

	if err != nil {
		t.Error(err)
	}

	if c.Client == "json" && len(c.Settings["json"]["url"]) == 0 {
		t.Error("Config for 'json' client has no 'url' field")
	}
}

func TestConfigLoadConfigEmptyString(t *testing.T) {
	_, err := LoadConfig("")

	if err == nil {
		t.Error("Expected error not found")
	}
}
