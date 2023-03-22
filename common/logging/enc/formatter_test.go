package enc

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestParseFormat(t *testing.T) {
	var tests = []struct {
		desc       string
		spec       string
		formatters []Formatter
	}{
		{
			desc:       "empty spec",
			spec:       "",
			formatters: []Formatter{},
		},
		{
			desc: "simple verb",
			spec: "%{color}",
			formatters: []Formatter{
				ColorFormatter{},
			},
		},
		{
			desc: "with prefix",
			spec: "prefix %{color}",
			formatters: []Formatter{
				StringFormatter{Value: "prefix "},
				ColorFormatter{},
			},
		},
		{
			desc: "with suffix",
			spec: "%{color} suffix",
			formatters: []Formatter{
				ColorFormatter{},
				StringFormatter{Value: " suffix"},
			},
		},
		{
			desc: "with prefix and suffix",
			spec: "prefix %{color} suffix",
			formatters: []Formatter{
				StringFormatter{Value: "prefix "},
				ColorFormatter{},
				StringFormatter{Value: " suffix"},
			},
		},
		{
			desc: "with format",
			spec: "%{level:.4s} suffix",
			formatters: []Formatter{
				LevelFormatter{FormatVerb: "%.4s"},
				StringFormatter{Value: " suffix"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			formatters, err := ParseFormat(test.spec)
			require.NoError(t, err)
			require.Equal(t, formatters, test.formatters)
		})
	}
}

func TestParseFormatError(t *testing.T) {
	_, err := ParseFormat("%{color:bad}")
	require.EqualError(t, err, "invalid color option: bad")
}

func TestNewFormatter(t *testing.T) {
	var tests = []struct {
		verb      string
		format    string
		formatter Formatter
		errorMsg  string
	}{
		{verb: "color", format: "", formatter: ColorFormatter{}},
		{verb: "color", format: "bold", formatter: ColorFormatter{Bold: true}},
		{verb: "color", format: "reset", formatter: ColorFormatter{Reset: true}},
		{verb: "color", format: "unknown", errorMsg: "invalid color option: unknown"},
		{verb: "id", format: "", formatter: SequenceFormatter{FormatVerb: "%d"}},
		{verb: "id", format: "04x", formatter: SequenceFormatter{FormatVerb: "%04x"}},
		{verb: "level", format: "", formatter: LevelFormatter{FormatVerb: "%s"}},
		{verb: "level", format: ".4s", formatter: LevelFormatter{FormatVerb: "%.4s"}},
		{verb: "message", format: "", formatter: MessageFormatter{FormatVerb: "%s"}},
		{verb: "message", format: "#30s", formatter: MessageFormatter{FormatVerb: "%#30s"}},
		{verb: "module", format: "", formatter: ModuleFormatter{FormatVerb: "%s"}},
		{verb: "module", format: "ok", formatter: ModuleFormatter{FormatVerb: "%ok"}},
		{verb: "shortfunc", format: "", formatter: ShortFuncFormatter{"%s"}},
		{verb: "shortfunc", format: "U", formatter: ShortFuncFormatter{FormatVerb: "%U"}},
		{verb: "time", format: "", formatter: TimeFormatter{Layout: "2006-01-02T15:04:05.999Z07:00"}},
		{verb: "time", format: time.RFC3339Nano, formatter: TimeFormatter{Layout: time.RFC3339Nano}},
		{verb: "unknown", format: "", errorMsg: "unknown verb: unknown"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test#%d", i), func(t *testing.T) {
			formatter, err := NewFormatter(test.verb, test.format)
			if err != nil {
				require.EqualError(t, err, test.errorMsg)
			} else {
				require.Equal(t, formatter, test.formatter)
			}
		})
	}
}

func TestColorFormatter(t *testing.T) {
	var tests = []struct {
		f         ColorFormatter
		level     zapcore.Level
		formatted string
	}{
		{f: ColorFormatter{Reset: true}, level: zapcore.DebugLevel, formatted: ResetColor()},
		{f: ColorFormatter{}, level: zapcore.DebugLevel, formatted: ColorCyan.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.DebugLevel, formatted: ColorCyan.Bold()},
		{f: ColorFormatter{}, level: zapcore.InfoLevel, formatted: ColorBlue.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.InfoLevel, formatted: ColorBlue.Bold()},
		{f: ColorFormatter{}, level: zapcore.WarnLevel, formatted: ColorYellow.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.WarnLevel, formatted: ColorYellow.Bold()},
		{f: ColorFormatter{}, level: zapcore.ErrorLevel, formatted: ColorRed.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.ErrorLevel, formatted: ColorRed.Bold()},
		{f: ColorFormatter{Reset: true}, level: zapcore.ErrorLevel, formatted: ResetColor()},
		{f: ColorFormatter{}, level: zapcore.DPanicLevel, formatted: ColorMagenta.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.DPanicLevel, formatted: ColorMagenta.Bold()},
		{f: ColorFormatter{Reset: true}, level: zapcore.DPanicLevel, formatted: ResetColor()},
		{f: ColorFormatter{}, level: zapcore.PanicLevel, formatted: ColorMagenta.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.PanicLevel, formatted: ColorMagenta.Bold()},
		{f: ColorFormatter{Reset: true}, level: zapcore.PanicLevel, formatted: ResetColor()},
		{f: ColorFormatter{}, level: zapcore.FatalLevel, formatted: ColorMagenta.Normal()},
		{f: ColorFormatter{Bold: true}, level: zapcore.FatalLevel, formatted: ColorMagenta.Bold()},
		{f: ColorFormatter{Reset: true}, level: zapcore.FatalLevel, formatted: ResetColor()},
		{f: ColorFormatter{Bold: true}, level: zapcore.Level(99), formatted: ColorNone.Bold()},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test#%d", i), func(t *testing.T) {
			buf := &bytes.Buffer{}
			entry := zapcore.Entry{Level: test.level}
			test.f.Format(buf, entry, nil)
			require.Equal(t, test.formatted, buf.String())
		})
	}
}
