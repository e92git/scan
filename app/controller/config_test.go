package controller

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := GetTestController()
	if err != nil {
		t.Error(err)
		return
	}
}

var TestController *Config
func GetTestController() (*Config, error) {
	if TestController != nil {
		return TestController, nil
	}
	c, err := NewTestEnv()
	if err != nil {
		return nil, err
	}
	TestController = c

	return c, nil
}