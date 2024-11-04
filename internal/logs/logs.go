package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type DefaultHook struct {
}

// 实现logrus的hook接口
func (h *DefaultHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel, log.FatalLevel, log.WarnLevel}
}

func (h *DefaultHook) Fire(entry *log.Entry, name string) error {
	entry.Data["appName"] = name
	file, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	line, _ := entry.String()
	file.Write([]byte(line))

	return nil
}
