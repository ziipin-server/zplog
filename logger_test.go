package zplog

import (
	"os"
	"testing"
	"fmt"
)

func TestSimpleLog(t *testing.T) {
	LogDebug("1")
	LogInfo("2")
	LogWarn("3")
	LogError("4")
	LogFatal("5")
}

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(LOG_WARN)
	LogDebug("1")
	LogInfo("2")
	LogWarn("3")
	LogError("4")
	LogFatal("5")
}

func TestNewLogger(t *testing.T) {
	SetLogLevel(LOG_WARN)
	LogDebug("1")
	LogInfo("2")
	LogWarn("%d", 3)
	LogError("4")
	LogFatal("5")
	logger := NewLogger(os.Stdout, "[logger]-")
	logger.SetLogLevel(LOG_DEBUG)
	logger.Debug("x %d", 1)
	logger.Info("x 2")
	logger.Warn("x 3")
	logger.Error("x 4")
	logger.Fatal("x 5")
}

type Room struct {
	Id int
}

func (r *Room)Work()  {
	r.LogError("something wrong here")
}

func (r *Room)LogError(format string,args... interface{})  {
	LogErrorUpper(1,fmt.Sprintf("room: %d %s",r.Id,fmt.Sprintf(format,args...)))
}

func TestLogUpper(t *testing.T)  {
	room := new(Room)
	room.Id = 10000
	room.Work()
}
