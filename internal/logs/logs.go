package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type DefaultHook struct {
}

type DefaultLog struct {
}

type AddLog interface {
	SetUpLogs()
}

// NewDefaultLog 返回一个DefaultLog
func NewDefaultLog() DefaultLog {
	return DefaultLog{}
}

// 实现logrus的hook接口
func (h *DefaultHook) Levels() []log.Level {
	return []log.Level{log.ErrorLevel, log.FatalLevel, log.WarnLevel}
}

func (h *DefaultHook) Fire(entry *log.Entry) error {
	entry.Data["appName"] = "ToDoList"
	file, err := os.OpenFile("internal/logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Error(err.Error())
	}
	defer file.Close()

	line, _ := entry.String()
	file.Write([]byte(line))
	return nil
}

func (dl DefaultLog) SetUpLogs() {
	log.AddHook(&DefaultHook{})
}
