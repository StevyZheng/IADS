package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = NewIadsLog()

func FatalPrintln(msg string) {
	if Logger.logger.Out != os.Stdout {
		println(msg)
	}
	Logger.LogFatal(msg)
	print("hahahahahaha")
}

type IadsLevel logrus.Level

const (
	InfoLevel  logrus.Level = logrus.InfoLevel
	WarnLevel  logrus.Level = logrus.WarnLevel
	ErrorLevel logrus.Level = logrus.ErrorLevel
	FatalLevel logrus.Level = logrus.FatalLevel
	DebugLevel logrus.Level = logrus.DebugLevel
)

type IadsLog struct {
	logger *logrus.Logger
}

func NewIadsLog() *IadsLog {
	iadsLog := new(IadsLog)
	iadsLog.logger = logrus.New()
	iadsLog.logger.Out = os.Stdout
	iadsLog.logger.Formatter = &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    true,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
	}
	iadsLog.logger.Level = logrus.InfoLevel
	return iadsLog
}

func (l *IadsLog) SetLogLevel(level IadsLevel) {
	l.logger.SetLevel((logrus.Level)(level))
}

func (l IadsLog) LogInfo(msg string) {
	l.logger.WithFields(logrus.Fields{}).Info(msg)
}

func (l IadsLog) LogFatal(msg string) {
	l.logger.WithFields(logrus.Fields{}).Fatal(msg)
}

func (l IadsLog) LogWarn(msg string) {
	l.logger.WithFields(logrus.Fields{}).Warn(msg)
}

func (l IadsLog) LogDebug(msg string) {
	l.logger.WithFields(logrus.Fields{}).Debug(msg)
}
func (l IadsLog) LogError(msg string) {
	l.logger.WithFields(logrus.Fields{}).Error(msg)
}
