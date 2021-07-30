package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/ueumd/logger/lfshook"
	"github.com/ueumd/logger/rotatelogs"
	"time"
)


var defaultLogger *logrus.Logger


func Init(filePath string, isStdout bool, loglevel uint32) error {
	if loglevel > 5 {
		loglevel = 5
	}

	logrus.StandardLogger().SetFormatter(&logrus.JSONFormatter{})
	defaultLogger = logrus.New()
	defaultLogger.SetLevel(logrus.Level(loglevel))
	defaultLogger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"})
	if !isStdout {
		defaultLogger.SetOutput(&emptyWriter{})
	}

	hk,err := newRotateHook(filePath, 7*24*time.Hour, 24*time.Hour)
	defaultLogger.AddHook(hk)
	return err
}


func newRotateHook(logPath string, maxAge time.Duration, rotationTime time.Duration) (*lfshook.LfsHook, error) {
	writer, err := rotatelogs.New(
		logPath + ".%Y%m%d",
		rotatelogs.WithLinkName(logPath),      // 生成软链，指向最新日志文
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)

	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", err)
		return nil, err
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.000"}), nil
}


//创建一个新的logger 实例
func NewLogger() *logrus.Logger {
	return logrus.New()
}

func WithFiled(key string, value interface{}) *logrus.Entry {
	return defaultLogger.WithField(key, value)
}

func WithFileds(fields map[string]interface{}) *logrus.Entry {
	return defaultLogger.WithFields(fields)
}

func Traceln(args ...interface{}) {
	defaultLogger.Traceln(args ...)
}

func TraceF(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v ...)
}

func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args ...)
}

func DebugF(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args ...)
}

func InfoF(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args ...)
}

func WarnF(format string, v ...interface{}) {
	defaultLogger.Warnf(format, v ...)
}

func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args ...)
}

func ErrorF(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args ...)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}






