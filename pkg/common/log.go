package common

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

var logger = logrus.New()

func init() {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "error"
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Panic("日志等级不合法！", err)
	}
	logger.Level = lvl
	logger.AddHook(ContextHook{})
}

type ContextHook struct {
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["file"] = path.Base(file)
		entry.Data["func"] = path.Base(funcName)
		entry.Data["line"] = line
	}
	return nil
}

func Logger() *logrus.Logger {
	return logger
}
