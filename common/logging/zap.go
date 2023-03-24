package logging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

type LarkLogger struct {
	sl *zap.SugaredLogger
}

func NewLarkLogger(zl *zap.Logger, options ...zap.Option) *LarkLogger {
	return &LarkLogger{sl: zl.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar()}
}

func (ll *LarkLogger) DPanic(args ...interface{})                    { ll.sl.DPanicf(formatArgs(args)) }
func (ll *LarkLogger) DPanicf(template string, args ...interface{})  { ll.sl.DPanicf(template, args...) }
func (ll *LarkLogger) DPanicw(msg string, kvPairs ...interface{})    { ll.sl.DPanicw(msg, kvPairs...) }
func (ll *LarkLogger) Debug(args ...interface{})                     { ll.sl.Debugf(formatArgs(args)) }
func (ll *LarkLogger) Debugf(template string, args ...interface{})   { ll.sl.Debugf(template, args...) }
func (ll *LarkLogger) Debugw(msg string, kvPairs ...interface{})     { ll.sl.Debugw(msg, kvPairs...) }
func (ll *LarkLogger) Error(args ...interface{})                     { ll.sl.Errorf(formatArgs(args)) }
func (ll *LarkLogger) Errorf(template string, args ...interface{})   { ll.sl.Errorf(template, args...) }
func (ll *LarkLogger) Errorw(msg string, kvPairs ...interface{})     { ll.sl.Errorw(msg, kvPairs...) }
func (ll *LarkLogger) Fatal(args ...interface{})                     { ll.sl.Fatalf(formatArgs(args)) }
func (ll *LarkLogger) Fatalf(template string, args ...interface{})   { ll.sl.Fatalf(template, args...) }
func (ll *LarkLogger) Fatalw(msg string, kvPairs ...interface{})     { ll.sl.Fatalw(msg, kvPairs...) }
func (ll *LarkLogger) Info(args ...interface{})                      { ll.sl.Infof(formatArgs(args)) }
func (ll *LarkLogger) Infof(template string, args ...interface{})    { ll.sl.Infof(template, args...) }
func (ll *LarkLogger) Infow(msg string, kvPairs ...interface{})      { ll.sl.Infow(msg, kvPairs...) }
func (ll *LarkLogger) Panic(args ...interface{})                     { ll.sl.Panicf(formatArgs(args)) }
func (ll *LarkLogger) Panicf(template string, args ...interface{})   { ll.sl.Panicf(template, args...) }
func (ll *LarkLogger) Panicw(msg string, kvPairs ...interface{})     { ll.sl.Panicw(msg, kvPairs...) }
func (ll *LarkLogger) Warn(args ...interface{})                      { ll.sl.Warnf(formatArgs(args)) }
func (ll *LarkLogger) Warnf(template string, args ...interface{})    { ll.sl.Warnf(template, args...) }
func (ll *LarkLogger) Warnw(msg string, kvPairs ...interface{})      { ll.sl.Warnw(msg, kvPairs...) }
func (ll *LarkLogger) Warning(args ...interface{})                   { ll.sl.Warnf(formatArgs(args)) }
func (ll *LarkLogger) Warningf(template string, args ...interface{}) { ll.sl.Warnf(template, args...) }

func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	opts := []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)}
	opts = append(opts, options...)
	return zap.New(core, opts...)
}

func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(zap.AddCaller(), zap.AddCallerSkip(3))
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

// formatArgs 仅仅是将args里的所有元素拼接在一起。
func formatArgs(args []interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(args...), "\n")
}