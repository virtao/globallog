package globallog

import (
	"bitbucket.org/kardianos/osext"
	"errors"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"time"
)

const (
	TAG             = "GlobalLog"
	LOG_FILE_NAME   = "log.log"
	LOG_FILE_FORMAT = "[%D %T] [%L] (%S) %M"
)

type GlobalLogger struct {
	tag string
	l4g.Logger
}

var _logInstance *GlobalLogger

func GetLogger() *GlobalLogger {
	if _logInstance == nil {
		initLogger()
	}
	return _logInstance
}

func initLogger() {
	_logInstance = &GlobalLogger{TAG, make(l4g.Logger, 2)}

	consoleLogWriter := l4g.NewConsoleLogWriter()
	fileLogWriter := l4g.NewFileLogWriter(getLogFilePath(), false)
	fileLogWriter.SetFormat(LOG_FILE_FORMAT)

	_logInstance.AddFilter("stdout", l4g.DEBUG, consoleLogWriter)
	_logInstance.AddFilter("file", l4g.DEBUG, fileLogWriter)
	_logInstance.Info(TAG, "日志模块已启动。")
}

func CloseLogger() {
	if _logInstance != nil {
		_logInstance.Info(TAG, "日志模块准备终止。")
		_logInstance.Close()
		_logInstance = nil
		time.Sleep(50 * time.Millisecond) //等待文件系统同步，防止日志保存不完整
	}
}

func ErrFormat(tag string, args ...interface{}) (err error) {
	return errors.New(fmt.Sprintf("[%s] %v", tag, args))
}

func getLogFilePath() (path string) {
	return getExeFilePath() + LOG_FILE_NAME
}

func getExeFilePath() (path string) {
	var err error
	if path, err = osext.ExecutableFolder(); err != nil {
		path = ""
	}
	return
}

func (l *GlobalLogger) Info(tag string, args ...interface{}) {
	l.Logger.Info("[%s] %v", tag, args)
}

func (l *GlobalLogger) Critical(tag string, args ...interface{}) (err error) {
	return l.Logger.Critical("[%s] %v", tag, args)
}

func (l *GlobalLogger) Error(tag string, args ...interface{}) (err error) {
	return l.Logger.Error("[%s] %v", tag, args)
}

func (l *GlobalLogger) Warn(tag string, args ...interface{}) (err error) {
	return l.Logger.Warn("[%s] %v", tag, args)
}

func (l *GlobalLogger) Debug(tag string, args ...interface{}) {
	l.Logger.Debug("[%s] %v", tag, args)
}

func (l *GlobalLogger) Fine(tag string, args ...interface{}) {
	l.Logger.Fine("[%s] %v", tag, args)
}

func (l *GlobalLogger) Finest(tag string, args ...interface{}) {
	l.Logger.Finest("[%s] %v", tag, args)
}
func (l *GlobalLogger) Trace(tag string, args ...interface{}) {
	l.Logger.Trace("[%s] %v", tag, args)
}
