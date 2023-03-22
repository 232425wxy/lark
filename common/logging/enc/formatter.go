package enc

import (
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"go.uber.org/zap/zapcore"
)

// color: 通过特定的颜色醒目的标示日志条目的级别；
// id: 唯一的日志条目序列号；
// level: 日志级别；
// message: 日志信息；
// module: zap 日志记录器的名称；
// shortfunc: 创建日志记录的函数的名称；
// time: 创建日志记录时的时间。
// color id level message module shortfunc time都是verb，冒号后面跟着的是format
var formatRegexp = regexp.MustCompile(`%{(color|id|level|message|module|shortfunc|time)(?::(.*?))?}`)

// sequence 是一个全局的序列号。
var sequence uint64

func SetSequence(s uint64) {
	atomic.StoreUint64(&sequence, s)
}

// ParseFormat 解析日志格式
func ParseFormat(spec string) ([]Formatter, error) {
	cursor := 0
	formatters := []Formatter{}

	matches := formatRegexp.FindAllStringSubmatchIndex(spec, -1)
	for _, m := range matches {
		start, end := m[0], m[1]
		verbStart, verbEnd := m[2], m[3]
		formatStart, formatEnd := m[4], m[5]

		if start > cursor {
			formatters = append(formatters, StringFormatter{Value: spec[cursor:start]})
		}
		var verb = spec[verbStart:verbEnd]
		var format string
		if formatStart >= 0 {
			format = spec[formatStart:formatEnd]
		}

		formatter, err := NewFormatter(verb, format)
		if err != nil {
			return nil, err
		}
		formatters = append(formatters, formatter)
		cursor = end
	}

	// 处理剩下的内容
	if cursor != len(spec) {
		formatters = append(formatters, StringFormatter{Value: spec[cursor:]})
	}
	return formatters, nil
}

func NewFormatter(verb, format string) (Formatter, error) {
	switch verb {
	case "color":
		return newColorFormatter(format)
	case "id":
		return newSequenceFormatter(format), nil
	case "level":
		return newLevelFormatter(format), nil
	case "message":
		return newMessageFormatter(format), nil
	case "module":
		return newModuleFormatter(format), nil
	case "shortfunc":
		return newShortFuncFormatter(format), nil
	case "time":
		return newTimeFormatter(format), nil
	default:
		return nil, fmt.Errorf("unknown verb: %s", verb)
	}
}

type MultiFormatter struct {
	mutex sync.RWMutex
	formatters []Formatter
}

func NewMultiFormatters(formatters ...Formatter) *MultiFormatter {
	return &MultiFormatter{
		formatters: formatters,
	}
}

func (mf *MultiFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	mf.mutex.RLock()
	for i := range mf.formatters {
		mf.formatters[i].Format(w, entry, fields)
	}
	mf.mutex.RUnlock()
}

func (mf *MultiFormatter) SetFormatters(formatters []Formatter) {
	mf.mutex.Lock()
	mf.formatters = formatters
	mf.mutex.Unlock()
}

type ColorFormatter struct {
	Bold bool // 设置粗体属性
	Reset bool // 重置颜色和属性，说白了就是清除之前添加的颜色和粗体效果
}

func (cf ColorFormatter) LevelColor(level zapcore.Level) Color {
	switch level {
	case zapcore.DebugLevel:
		return ColorCyan 
	case zapcore.InfoLevel:
		return ColorBlue
	case zapcore.WarnLevel:
		return ColorYellow
	case zapcore.ErrorLevel:
		return ColorRed
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		return ColorMagenta
	default:
		return ColorNone
	}
}

func (cf ColorFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	switch {
	case cf.Reset:
		fmt.Fprint(w, ResetColor())
	case cf.Bold:
		fmt.Fprint(w, cf.LevelColor(entry.Level).Bold())
	default:
		fmt.Fprint(w, cf.LevelColor(entry.Level).Normal())
	}
}

func newColorFormatter(f string) (ColorFormatter, error) {
	switch f {
	case "bold":
		return ColorFormatter{Bold: true}, nil
	case "reset":
		return ColorFormatter{Reset: true}, nil
	case "":
		return ColorFormatter{}, nil
	default:
		return ColorFormatter{}, fmt.Errorf("invalid color option: %s", f)
	}
}

type LevelFormatter struct {
	FormatVerb string // 就是格式化样式，所谓动词的含义，就是规定格式，输出的内容必须按照这个格式输出，例如："age:%d"
}

// Format 日志级别会以大写字母的样式输出
func (lf LevelFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, lf.FormatVerb, entry.Level.CapitalString())
}

func newLevelFormatter(f string) LevelFormatter {
	return LevelFormatter{FormatVerb: "%" + stringOrDefault(f, "s")}
}

type MessageFormatter struct {
	FormatVerb string
}

func (mf MessageFormatter) Format(w io.Writer, entery zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, mf.FormatVerb, strings.TrimRight(entery.Message, "\n")) // 清除掉message右边的所有换行符
}

func newMessageFormatter(f string) MessageFormatter {
	return MessageFormatter{FormatVerb: "%" + stringOrDefault(f, "s")}
}

// ModuleFormatter module实际上就是日志记录器的名字
type ModuleFormatter struct {
	FormatVerb string
}

func (mf ModuleFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, mf.FormatVerb, entry.LoggerName)
}

func newModuleFormatter(f string) ModuleFormatter {
	return ModuleFormatter{FormatVerb: "%" + stringOrDefault(f, "s")}
}

type SequenceFormatter struct {
	FormatVerb string
}

func (sf SequenceFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, sf.FormatVerb, atomic.AddUint64(&sequence, 1))
}

func newSequenceFormatter(f string) SequenceFormatter {
	return SequenceFormatter{FormatVerb: "%" + stringOrDefault(f, "d")}
}

type ShortFuncFormatter struct {
	FormatVerb string
}

func (sff ShortFuncFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	f := runtime.FuncForPC(entry.Caller.PC)
	if f== nil {
		fmt.Fprintf(w, sff.FormatVerb, "unknown")
		return
	}
	fname := f.Name()
	funcIdx := strings.LastIndex(fname, ".")
	fmt.Fprintf(w, sff.FormatVerb, fname[funcIdx+1:])
}

func newShortFuncFormatter(f string) ShortFuncFormatter {
	return ShortFuncFormatter{FormatVerb: "%" + stringOrDefault(f, "s")}
}

type TimeFormatter struct {
	Layout string
}

func (tf TimeFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprint(w, entry.Time.Format(tf.Layout))
}

func newTimeFormatter(f string) TimeFormatter {
	return TimeFormatter{Layout: stringOrDefault(f, "2006-01-02T15:04:05.999Z07:00")}
}

// stringOrDefault 如果第一个参数不为""，则返回第一个参数，否则返回第二个参数。
func stringOrDefault(str, defaultStr string) string {
	if str != "" {
		return str
	}
	return defaultStr
}