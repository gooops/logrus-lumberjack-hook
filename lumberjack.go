package logrus_lumberjack

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type LumberjackHook struct {
	Writer *lumberjack.Logger
}

func (hook *LumberjackHook) Write(line string) error {
	_, err := hook.Writer.Write([]byte(line))
	return err
}

func NewLumberjackHook(lumberConfig *lumberjack.Logger) (*LumberjackHook, error) {
	return &LumberjackHook{Writer: lumberConfig}, nil
}

func (hook *LumberjackHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Write(line)
	case logrus.FatalLevel:
		return hook.Write(line)
	case logrus.ErrorLevel:
		return hook.Write(line)
	case logrus.WarnLevel:
		return hook.Write(line)
	case logrus.InfoLevel:
		return hook.Write(line)
	case logrus.DebugLevel:
		return hook.Write(line)
	default:
		return nil
	}
}

func (hook *LumberjackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
