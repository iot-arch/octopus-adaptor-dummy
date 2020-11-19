package zap

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap"
)

// nullInfoLogger doesn't log any log.
type nullInfoLogger struct{}

func (nullInfoLogger) Enabled() bool                                             { return false }
func (nullInfoLogger) Info(_ string, _ ...interface{})                           {}
func (l nullInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (nullInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l nullInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l nullInfoLogger) WithName(name string) logr.Logger                        { return l }

// DebugInfoLogger logs as Debug level.
type debugInfoLogger struct {
	z *zap.Logger
}

func (l debugInfoLogger) Enabled() bool { return true }
func (l debugInfoLogger) Info(msg string, keysAndValues ...interface{}) {
	l.z.Debug(msg, handleFields(keysAndValues)...)
}
func (l debugInfoLogger) ToZapLogger() *zap.Logger                                { return l.z }
func (l debugInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (debugInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l debugInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l debugInfoLogger) WithName(name string) logr.Logger                        { return l }

// infoInfoLogger logs as Info level.
type infoInfoLogger struct {
	z *zap.Logger
}

func (l infoInfoLogger) Enabled() bool { return true }
func (l infoInfoLogger) Info(msg string, keysAndValues ...interface{}) {
	l.z.Info(msg, handleFields(keysAndValues)...)
}
func (l infoInfoLogger) ToZapLogger() *zap.Logger                                { return l.z }
func (l infoInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (infoInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l infoInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l infoInfoLogger) WithName(name string) logr.Logger                        { return l }

// warnInfoLogger logs as Warn level.
type warnInfoLogger struct {
	z *zap.Logger
}

func (l warnInfoLogger) Enabled() bool { return true }
func (l warnInfoLogger) Info(msg string, keysAndValues ...interface{}) {
	l.z.Warn(msg, handleFields(keysAndValues)...)
}
func (l warnInfoLogger) ToZapLogger() *zap.Logger                                { return l.z }
func (l warnInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (warnInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l warnInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l warnInfoLogger) WithName(name string) logr.Logger                        { return l }

// errorInfoLogger logs as Error level.
type errorInfoLogger struct {
	z *zap.Logger
}

func (l errorInfoLogger) Enabled() bool { return true }
func (l errorInfoLogger) Info(msg string, keysAndValues ...interface{}) {
	l.z.Error(msg, handleFields(keysAndValues)...)
}
func (l errorInfoLogger) ToZapLogger() *zap.Logger { return l.z }

func (l errorInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (errorInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l errorInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l errorInfoLogger) WithName(name string) logr.Logger                        { return l }

// fatalInfoLogger logs as Fatal level.
type fatalInfoLogger struct {
	z *zap.Logger
}

func (l fatalInfoLogger) Enabled() bool { return true }
func (l fatalInfoLogger) Info(msg string, keysAndValues ...interface{}) {
	l.z.Fatal(msg, handleFields(keysAndValues)...)
}
func (l fatalInfoLogger) ToZapLogger() *zap.Logger { return l.z }

func (l fatalInfoLogger) V(level int) logr.InfoLogger                             { return l }
func (fatalInfoLogger) Error(err error, msg string, keysAndValues ...interface{}) {}
func (l fatalInfoLogger) WithValues(keysAndValues ...interface{}) logr.Logger     { return l }
func (l fatalInfoLogger) WithName(name string) logr.Logger                        { return l }

// NewNullInfoLogr returns a null logr.InfoLogger.
func NewNullInfoLogr() logr.InfoLogger {
	return nullInfoLogger{}
}

// WrapAsDebugInfoLogr wraps a Zap logger as a logr.InfoLogger to logs in Debug level.
func WrapAsDebugInfoLogr(z *zap.Logger) logr.InfoLogger {
	return debugInfoLogger{z: z}
}

// WrapAsInfoInfoLogr wraps a Zap logger as a logr.InfoLogger to logs in Info level.
func WrapAsInfoInfoLogr(z *zap.Logger) logr.InfoLogger {
	return infoInfoLogger{z: z}
}

// WrapAsWarnInfoLogr wraps a Zap logger as a logr.InfoLogger to logs in Warn level.
func WrapAsWarnInfoLogr(z *zap.Logger) logr.InfoLogger {
	return warnInfoLogger{z: z}
}

// WrapAsErrorInfoLogr wraps a Zap logger as a logr.InfoLogger to logs in Error level.
func WrapAsErrorInfoLogr(z *zap.Logger) logr.InfoLogger {
	return errorInfoLogger{z: z}
}

// WrapAsFatalInfoLogr wraps a Zap logger as a logr.InfoLogger to logs in Fatal level.
func WrapAsFatalInfoLogr(z *zap.Logger) logr.InfoLogger {
	return fatalInfoLogger{z: z}
}
