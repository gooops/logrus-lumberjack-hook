package logrus_lumberjack

import (
	"errors"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type LumberjackHook struct {
	Crit  *lumberjack.Logger
	Err   *lumberjack.Logger
	Warn  *lumberjack.Logger
	Info  *lumberjack.Logger
	Debug *lumberjack.Logger
}

const (
	DEFAULTMAXSIZE   = 1 // megabytes
	DEFAULTMAXBACKUP = 3
	DEFAULTMAXAGE    = 30 //days
)

func NewLumberjackHook(lumber interface{}) (*LumberjackHook, error) {
	defConfig := &lumberjack.Logger{
		Filename:   "stdout.log",
		MaxSize:    DEFAULTMAXSIZE, // megabytes
		MaxBackups: DEFAULTMAXBACKUP,
		MaxAge:     DEFAULTMAXAGE, //days
	}
	lumberhook := &LumberjackHook{
		Crit:  defConfig,
		Err:   defConfig,
		Warn:  defConfig,
		Info:  defConfig,
		Debug: defConfig,
	}

	switch m := lumber.(type) {
	case map[string]*lumberjack.Logger:
		for k, v := range m {
			switch k {
			case "crit":
				lumberhook.Crit = v
			case "err":
				lumberhook.Err = v
			case "warn":
				lumberhook.Warn = v
			case "info":
				lumberhook.Info = v
			case "debug":
				lumberhook.Debug = v
			default:
				return nil, errors.New(k + " does not support!")
			}
		}
	case map[string]string:
		for k, v := range m {
			switch k {
			case "crit":
				lumberhook.Crit = &lumberjack.Logger{
					Filename:   v,
					MaxSize:    DEFAULTMAXSIZE, // megabytes
					MaxBackups: DEFAULTMAXBACKUP,
					MaxAge:     DEFAULTMAXAGE, //days
				}
			case "err":
				lumberhook.Err = &lumberjack.Logger{
					Filename:   v,
					MaxSize:    DEFAULTMAXSIZE, // megabytes
					MaxBackups: DEFAULTMAXBACKUP,
					MaxAge:     DEFAULTMAXAGE, //days
				}
			case "warn":
				lumberhook.Warn = &lumberjack.Logger{
					Filename:   v,
					MaxSize:    DEFAULTMAXSIZE, // megabytes
					MaxBackups: DEFAULTMAXBACKUP,
					MaxAge:     DEFAULTMAXAGE, //days
				}
			case "info":
				lumberhook.Info = &lumberjack.Logger{
					Filename:   v,
					MaxSize:    DEFAULTMAXSIZE, // megabytes
					MaxBackups: DEFAULTMAXBACKUP,
					MaxAge:     DEFAULTMAXAGE, //days
				}
			case "debug":
				lumberhook.Debug = &lumberjack.Logger{
					Filename:   v,
					MaxSize:    DEFAULTMAXSIZE, // megabytes
					MaxBackups: DEFAULTMAXBACKUP,
					MaxAge:     DEFAULTMAXAGE, //days
				}
			default:
				return nil, errors.New(k + " does not support!")
			}
		}

	default:
		fmt.Printf("%T", lumber)
	}
	fmt.Println(lumberhook)
	return lumberhook, nil
}

func (hook *LumberjackHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	switch entry.Level {
	case logrus.PanicLevel:
		_, err = hook.Crit.Write([]byte(line))
	case logrus.FatalLevel:
		_, err = hook.Crit.Write([]byte(line))
	case logrus.ErrorLevel:
		_, err = hook.Err.Write([]byte(line))
	case logrus.WarnLevel:
		_, err = hook.Warn.Write([]byte(line))
	case logrus.InfoLevel:
		_, err = hook.Info.Write([]byte(line))
	case logrus.DebugLevel:
		_, err = hook.Debug.Write([]byte(line))
	default:
		return nil
	}
	return err
}

func (hook *LumberjackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
