package log

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	logrus_lumberjack "github.com/fallais/logrus-lumberjack-hook"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defaultFormatter = &Formatter{
	FieldsOrder:     []string{"request_id", "http_code", "total_time", "ip", "method", "uri", "err_code", "err_msg"},
	TimestampFormat: time.RFC3339,
	CallerFirst:     true,
	NoFieldsBracket: true,
	CustomCallerFormatter: func(f *runtime.Frame) string {
		_, filename := filepath.Split(f.File)
		return fmt.Sprintf("%18s:%-4d", filename, f.Line)
	},
}

func New(level string, filePath string, fileName string, maxSize int64, maxAge time.Duration, rotationTime time.Duration) *logrus.Logger {
	log := logrus.New()
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		panic(err)
	}

	log.SetFormatter(defaultFormatter)
	log.SetReportCaller(true)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	log.SetLevel(lvl)
	log.SetOutput(os.Stdout)

	log.AddHook(newRotateHook(filePath, fileName, int(maxSize), 0))
	return log
}

func newRotateHook(logPath string, logFileName string, maxSize int, maxAge int) logrus.Hook {
	baseLogName := path.Join(logPath, logFileName)

	// Set the Lumberjack logger
	lumberjackLogger := &lumberjack.Logger{
		Filename:   baseLogName,
		MaxSize:    maxSize,
		MaxBackups: 3,
		MaxAge:     maxAge,
		LocalTime:  true,
	}

	// Add Lumberjack hook
	hook, err := logrus_lumberjack.NewLumberjackHook(lumberjackLogger)
	if err != nil {
		logrus.Fatalln("Unable to add the Lumberjack hook :", err)
	}

	return hook
}
