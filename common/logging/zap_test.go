package logging

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/232425wxy/lark/common/logging/enc"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFmtSprintln(t *testing.T) {
	args := []interface{}{1, "一", 2, "二"}
	s := fmt.Sprintln(args...)
	t.Log(strings.TrimSuffix(s, "\n"))
}

func TestLarkLoggerEncoding(t *testing.T) {
	formatters, err := enc.ParseFormat("%{color}[%{module}] %{shortfunc} -> %{level:.4s}%{color:reset} %{message}")
	require.NoError(t, err)
	encoder := enc.NewFormatEncoder(formatters...)

	buf := &bytes.Buffer{}
	core := zapcore.NewCore(encoder, zapcore.AddSync(buf), zap.NewAtomicLevel())
	zl := NewZapLogger(core).Named("test").With(zap.String("extra", "field"))
	fl := NewLarkLogger(zl)

	buf.Reset()
	fl.Info("string value", 0, 1.23, struct{}{})
	require.Equal(t, "\x1b[34m[test] TestLarkLoggerEncoding -> INFO\x1b[0m string value 0 1.23 {} extra=field\n", buf.String())
	
	buf.Reset()
	fl.Infof("string %s, %d, %.3f, %v", "strval", 0, 1.23, struct{}{})
	require.Equal(t, "\x1b[34m[test] TestLarkLoggerEncoding -> INFO\x1b[0m string strval, 0, 1.230, {} extra=field\n", buf.String())

	buf.Reset()
	fl.Infow("this is a message", "int", 0, "float", 1.23, "struct", struct{}{})
	require.Equal(t, "\x1b[34m[test] TestLarkLoggerEncoding -> INFO\x1b[0m this is a message extra=field int=0 float=1.23 struct={}\n", buf.String())
}