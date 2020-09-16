package config

import (
	"testing"
	"util/logger"
)

func TestGetConfigPath(t *testing.T) {
	path := GetConfigPath()
	logger.Info.Println(path)
}
