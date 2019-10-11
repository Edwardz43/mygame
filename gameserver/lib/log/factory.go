package log

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Create ...
func Create(t string) *logrus.Logger {
	logger := logrus.New()
	filename := time.Now().Format("20060102")
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/%v/%v/foo.log", t, filename),
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	logger.SetReportCaller(true)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
