# logrus-lumberjack-hook
```go
package main

import (
	log "github.com/Sirupsen/logrus"
	lumberhook "github.com/gooops/logrus-lumberjack-hook"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
}

func main() {
	// map指定不同日志级别的详细配置，需要导入 lumberjack "gopkg.in/natefinch/lumberjack.v2" 包
	// lhook, _ := lumberhook.NewLumberjackHook(map[string]*lumberjack.Logger{
	// 	"info": &lumberjack.Logger{
	// 		Filename:   "/var/log/foo.log",
	// 		MaxSize:    1, // megabytes
	// 		MaxBackups: 3,
	// 		MaxAge:     28, //days
	// 	}})

	// 只写日志文件，其他选项默认
	lhook, _ := lumberhook.NewLumberjackHook(map[string]string{
		"info": "/var/log/info.log",
		"warn": "/var/log/warn.log",
	})
	log.AddHook(lhook)
	log.Errorln("error log")
	log.Infoln("info log")
	log.Warningln("warn log")
}

```
