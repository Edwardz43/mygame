package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/Edwardz43/logrustash"
	"github.com/Edwardz43/mygame/gameserver/config"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger ...
type Logger struct {
	log    *logrus.Logger
	target string
}

// Printf print format string with info level log.
func (l *Logger) Printf(format string, args ...interface{}) {
	if pc, f, line, ok := runtime.Caller(1); ok {
		fnName := strings.Split(runtime.FuncForPC(pc).Name(), "gameserver")[1]
		// log.Printf("FILENAME : %v", fnName)
		// log.Printf("f : %v", f)
		file := strings.Split(f, "gameserver")[1]
		caller := fmt.Sprintf("%s:%v %s", file, line, fnName)
		l.log.WithField("caller", caller).Info(fmt.Sprintf(format, args...))
	}
}

// Println print string with info level log.
func (l *Logger) Println(msg interface{}) {
	if pc, f, line, ok := runtime.Caller(1); ok {
		fnName := strings.Split(runtime.FuncForPC(pc).Name(), "gameserver")[1]
		// log.Printf("FILENAME : %v", fnName)
		// log.Printf("f : %v", f)
		file := strings.Split(f, "gameserver")[1]
		caller := fmt.Sprintf("%s:%v %s", file, line, fnName)
		l.log.WithField("caller", caller).Info(fmt.Sprintf("%v", msg))
	}
}

func (l *Logger) Error(msg interface{}) {
	// TODO
	// l.log.SetOutput(&lumberjack.Logger{
	// 	Filename: fmt.Sprintf("logs/%v/%v/foo.log", l.target, filename),
	// })
}

// Create creates logrus Logger with specific target class.
func Create(t string) *Logger {
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
	logger.SetLevel(logrus.DebugLevel)
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	if config.GetELKConfig() == true {
		url := config.GetLogstashConfig()

		hook, err := logrustash.NewHook("tcp", url, t)
		if err != nil {
			fmt.Println(err)
		}
		logger.Hooks.Add(hook)
	}

	return &Logger{
		log:    logger,
		target: t,
	}
}
