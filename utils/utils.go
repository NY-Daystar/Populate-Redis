package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/NY-Daystar/populate-redis/configuration"
	"github.com/Pallinder/sillyname-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// GenerateName generate fake string similate to name
func GenerateName() string {
	name := sillyname.GenerateStupidName()
	return strings.ReplaceAll(name, " ", "-")
}

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// Source : https://gist.github.com/r0l1/3dcbb0c8f6cfe9c66ab8008f55f8f28b
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func AskForConfirmation() bool {
	var input = ReadValue()

	input = strings.ToLower(strings.TrimSpace(input))

	if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
		return true
	}

	return false
}

// ReadValue read input from terminal and returns its value
func ReadValue() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// setLogger create logger with zap librairy
func SetLogger(verbose *bool) *zap.Logger {
	// Setup log level
	level := zapcore.InfoLevel
	if *verbose {
		level = zap.DebugLevel
	}

	// Create logs folder if not exists
	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// Creating file backup
	lLogger := lumberjack.Logger{
		Filename:   configuration.LOGS_PATH, // File path
		MaxSize:    500,                     // 500 megabytes per files
		MaxBackups: 3,                       // 3 files before rotate
		MaxAge:     15,                      // 15 days
	}

	fileCfg := zap.NewProductionEncoderConfig()
	fileCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileCfg)

	// Specify encoding for logger
	var encoderTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("[2024-03-03 15:16:05]"))
	}

	concoleCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "message",
		EncodeTime:     encoderTime,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	consoleEncoder := zapcore.NewConsoleEncoder(concoleCfg)

	// Include output
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(&lLogger), zapcore.DebugLevel),
	)
	var logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Assurez-vous de vider les buffers avant de quitter
	logger.Sync()

	logger.Debug("Zap logger set",
		zap.String("filename", lLogger.Filename),
		zap.Int("filesize", lLogger.MaxSize),
		zap.Int("backupfile", lLogger.MaxBackups),
		zap.Int("fileage", lLogger.MaxAge),
	)

	return logger
}
