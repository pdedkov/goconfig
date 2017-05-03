package goconfig

import "testing"

// Config app config
type Config struct {
	StringVar string
	BoolVar   bool
	IntVal    int
}

// FakeConfig struct with invali field
type FakeConfig struct {
	BoolVar bool
}

func TestShouldCheckErrors(t *testing.T) {
	var c Config
	if err := NewConfigFromFile("", &c); err == nil {
		t.Error("Wrong response. Empty pass")
	}
	if err := NewConfigFromFile("./undefined", &c); err == nil {
		t.Error("Wrong response. Invalid file")
	}
	var f FakeConfig
	if err := NewConfigFromFile("", &f); err == nil {
		t.Error("Wrong response. Invalid struct")
	}

}

func TestNewConfigFromFile(t *testing.T) {
	var c Config
	if err := NewConfigFromFile("./config.toml", &c); err != nil {
		t.Error(err)
	}
	if !c.BoolVar {
		t.Error("Invalid bool var")
	}
}
