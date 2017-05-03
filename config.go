package goconfig

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

// NewConfigFromFile load config from file
func NewConfigFromFile(path string, conf interface{}) error {
	if path == "" {
		return errors.New("Empty path")
	}
	// check config file
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("File '%s' does not exists", path)
		}
		return err
	}

	if _, err := toml.DecodeFile(path, conf); err != nil {
		return fmt.Errorf("Error while loading config file: %v", err)
	}

	return nil
}
