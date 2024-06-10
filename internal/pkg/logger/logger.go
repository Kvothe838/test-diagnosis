package logger

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

const formatWithMilliseconds = "2006-01-02T15:04:05.999Z07:00"

var (
	once          sync.Once
	defaultLogger *logrus.Logger
	contextKeys   = []ContextKey{
		PatientID,
	}
)

type ContextKey string

const (
	PatientID ContextKey = "patient_id"
)

func buildLogger() *logrus.Logger {
	once.Do(func() {
		defaultLogger = logrus.New()
		defaultLogger.SetOutput(os.Stdout)
		defaultLogger.Out = os.Stdout

		logLevel := buildLogLevel()
		defaultLogger.SetLevel(logLevel)

		formatter := buildFormatter()
		defaultLogger.SetFormatter(formatter)
	})

	return defaultLogger
}

func buildLogLevel() logrus.Level {
	logLevel := logrus.InfoLevel
	if val, ok := os.LookupEnv("LOG_LEVEL"); ok {
		parsedLevel, err := logrus.ParseLevel(val)
		if err == nil {
			logLevel = parsedLevel
		}
	}

	return logLevel
}

func buildFormatter() logrus.Formatter {
	var formatter logrus.Formatter
	formatter = &logrus.JSONFormatter{
		TimestampFormat: formatWithMilliseconds,
	}
	if val, ok := os.LookupEnv("LOG_FORMATTER"); ok {
		if strings.EqualFold(val, "text") {
			formatter = &logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: formatWithMilliseconds,
			}
		}
	}
	return formatter
}

func SetPatientID(ctx context.Context, patientID string) context.Context {
	return context.WithValue(ctx, PatientID, patientID)
}

func loggerWithCtx(ctx context.Context) *logrus.Entry {
	l := buildLogger()
	entry := l.WithContext(ctx)
	for _, key := range contextKeys {
		if val := ctx.Value(key); val != nil && fmt.Sprint(val) != "" {
			entry = entry.WithField(string(key), val)
		}
	}

	return entry
}

func CtxDebug(ctx context.Context, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Debugln(args...)
}

func CtxDebugf(ctx context.Context, format string, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Debugf(format, args...)
}

func CtxInfo(ctx context.Context, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Infoln(args...)
}

func CtxInfof(ctx context.Context, format string, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Infof(format, args...)
}

func CtxError(ctx context.Context, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Errorln(args...)
}

func CtxErrorf(ctx context.Context, format string, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Errorf(format, args...)
}

func CtxFatal(ctx context.Context, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Fatal(args...)
}

func CtxWarn(ctx context.Context, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Warn(args...)
}

func CtxWarnf(ctx context.Context, format string, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwc.Warnf(format, args...)
}

func CtxFieldsInfo(ctx context.Context, fields logrus.Fields, args ...interface{}) {
	lwc := loggerWithCtx(ctx)
	lwf := lwc.WithFields(fields)
	lwf.Info(args...)
}
