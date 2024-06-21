package logging

//日志
import (
	"chat-system/config/global"
	"github.com/sirupsen/logrus"
)

// Setup 初始化日志配置
func Setup() {
	global.Config.Logger = logrus.New()
}

// Debug 调试
func Debug(v ...interface{}) {
	global.Config.Logger.Debug(v)
}

func Debugln(v ...interface{}) {
	global.Config.Logger.Debugln(v)
}

func Debugf(format string, args ...interface{}) {
	global.Config.Logger.Debugf(format, args...)
}

// Info 正常级别日志
func Info(v ...interface{}) {
	global.Config.Logger.Info(v)
}

func Infoln(v ...interface{}) {
	global.Config.Logger.Infoln(v)
}

func Infof(format string, args ...interface{}) {
	global.Config.Logger.Infof(format, args...)
}

// Warn 警告级别日志
func Warn(v ...interface{}) {
	global.Config.Logger.Warn(v)
}

func Warnln(v ...interface{}) {
	global.Config.Logger.Warnln(v)
}

func Warnf(format string, args ...interface{}) {
	global.Config.Logger.Warnf(format, args...)
}

// Error 错误级别日志
func Error(v ...interface{}) {
	global.Config.Logger.Error(v)
}

func Errorln(v ...interface{}) {
	global.Config.Logger.Errorln(v)
}

func Errorf(format string, args ...interface{}) {
	global.Config.Logger.Errorf(format, args...)
}

// Fatal 致命错误级别日志
func Fatal(v ...interface{}) {
	global.Config.Logger.Fatal(v)
}

func Fatalln(v ...interface{}) {
	global.Config.Logger.Fatalln(v)
}

func Fatalf(format string, args ...interface{}) {
	global.Config.Logger.Fatalf(format, args...)
}
