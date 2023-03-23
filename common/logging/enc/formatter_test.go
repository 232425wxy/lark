package enc

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
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

func TestLevelFormatter(t *testing.T) {
	var tests = []struct {
		level     zapcore.Level
		formatted string
	}{
		{level: zapcore.DebugLevel, formatted: "DEBUG"},
		{level: zapcore.InfoLevel, formatted: "INFO"},
		{level: zapcore.WarnLevel, formatted: "WARN"},
		{level: zapcore.ErrorLevel, formatted: "ERROR"},
		{level: zapcore.DPanicLevel, formatted: "DPANIC"},
		{level: zapcore.PanicLevel, formatted: "PANIC"},
		{level: zapcore.FatalLevel, formatted: "FATAL"},
		{level: zapcore.Level(99), formatted: "LEVEL(99)"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test#%d", i), func(t *testing.T) {
			buf := &bytes.Buffer{}
			entry := zapcore.Entry{Level: test.level}
			LevelFormatter{FormatVerb: "%s"}.Format(buf, entry, nil)
			require.Equal(t, buf.String(), test.formatted)
		})
	}
}

func TestMessageFormatter(t *testing.T) {
	buf := &bytes.Buffer{}
	entry := zapcore.Entry{Message: "some message text \n\n"}
	f := MessageFormatter{FormatVerb: "%s"}
	f.Format(buf, entry, nil)
	require.Equal(t, buf.String(), "some message text ")
}

func TestModuleFormatter(t *testing.T) {
	buf := &bytes.Buffer{}
	entry := zapcore.Entry{LoggerName: "logger/name"}
	f := ModuleFormatter{FormatVerb: "%s"}
	f.Format(buf, entry, nil)
	require.Equal(t, "logger/name", buf.String())
}

func TestSequenceFormatter(t *testing.T) {
	mutex := &sync.Mutex{}
	results := map[string]struct{}{}

	ready := &sync.WaitGroup{}
	ready.Add(100)

	finished := &sync.WaitGroup{}
	finished.Add(100)

	SetSequence(0)

	for i := 1; i <= 100; i++ {
		go func(i int) {
			buf := &bytes.Buffer{}
			entry := zapcore.Entry{Level: zapcore.DebugLevel}
			f := SequenceFormatter{FormatVerb: "%d"}
			ready.Done() // setup complete
			ready.Wait() // wait for all go routines to be ready

			f.Format(buf, entry, nil) // format concurrently

			mutex.Lock()
			results[buf.String()] = struct{}{}
			mutex.Unlock()

			finished.Done()
		}(i)
	}

	finished.Wait()
	for i := 1; i <= 100; i++ {
		require.Contains(t, results, strconv.Itoa(i))
	}
}

func TestShortFuncFormatter(t *testing.T) {
	callerpc, _, _, ok := runtime.Caller(0)
	require.True(t, ok)
	buf := &bytes.Buffer{}
	entry := zapcore.Entry{Caller: zapcore.EntryCaller{PC: callerpc}}
	ShortFuncFormatter{FormatVerb: "%s"}.Format(buf, entry, nil)
	require.Equal(t, "TestShortFuncFormatter", buf.String())

	buf.Reset()
	entry = zapcore.Entry{Caller: zapcore.EntryCaller{PC: 0}}
	ShortFuncFormatter{FormatVerb: "%s"}.Format(buf, entry, nil)
	require.Equal(t, "unknown", buf.String())
}

func TestTimeFormatter(t *testing.T) {
	buf := &bytes.Buffer{}
	entry := zapcore.Entry{Time: time.Date(2023, time.March, 23, 9, 40, 59, 333, time.Local)}
	f := TimeFormatter{Layout: time.RFC3339Nano}
	f.Format(buf, entry, nil)
	require.Equal(t, "2023-03-23T09:40:59.000000333+08:00", buf.String())
}

func TestMultiFormatter(t *testing.T) {
	entry := zapcore.Entry{
		Message: "message",
		Level:   zapcore.InfoLevel,
	}
	fields := []zapcore.Field{zap.String("key", "value")}

	var tests = []struct{
		desc string
		initial []Formatter
		update []Formatter
		expected string
	}{
		{
			desc: "no formatters",
			initial: nil,
			update: nil,
			expected: "",
		},
		{
			desc: "initial formatters",
			initial: []Formatter{StringFormatter{Value: "string1"}},
			update: nil,
			expected: "string1",
		},
		{
			desc: "set to formatters",
			initial: []Formatter{StringFormatter{Value: "string1"}},
			update: []Formatter{
				StringFormatter{Value: "string1"},
				StringFormatter{Value: "-"},
				StringFormatter{Value: "string2"},
			},
			expected: "string1-string2",
		},
		{
			desc: "set to empty",
			initial: []Formatter{StringFormatter{Value: "string1"}},
			update: []Formatter{},
			expected: "",
		},
	}

	for _, test := range tests {
		mf := NewMultiFormatters(test.initial...)
		if test.update != nil {
			mf.SetFormatters(test.update)
		}
		buf := &bytes.Buffer{}
		mf.Format(buf, entry, fields)
		require.Equal(t, test.expected, buf.String())
	}
}
