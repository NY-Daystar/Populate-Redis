package configuration

import (
	"encoding/json"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

const (
	PROJECT_NAME    string = "populate-redis"
	PROJECT_VERSION string = "1.0.0"
	LOGS_PATH       string = "logs/log.json" // Path of the logs
)

// Config: Configuration of the application
type Config struct {
	Redis RedisConfig `json:"redis"`
}

// RedisConfig : Configuration of redis to get a connection to a redis database
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int64  `json:"database"`
}

// ReadConfig read json file and convert into config struct
func ReadConfig(cfg *Config, logger *zap.Logger) {
	configFileName := "config.json"
	configFileName, _ = filepath.Abs(configFileName)
	logger.Sugar().Debugf("Loading config: %s", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		logger.Sugar().Fatalf("File error: %v", err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		logger.Sugar().Fatalf("Config error: %v", err.Error())
	}
}
