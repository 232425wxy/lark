package enc

import (
	"fmt"
	"io"
	"time"

	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type Formatter interface {
	Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field)
}

type FormatEncoder struct {
	zapcore.Encoder
	formatters []Formatter
	pool       buffer.Pool
}

func NewFormatEncoder(formatters ...Formatter) *FormatEncoder {
	return &FormatEncoder{
		Encoder: zaplogfmt.NewEncoder(
			zapcore.EncoderConfig{
				MessageKey:     "",
				LevelKey:       "",
				TimeKey:        "",
				NameKey:        "",
				CallerKey:      "",
				StacktraceKey:  "",
				LineEnding:     "\n",
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.Format("2006-01-02T15:04:05.999Z07:00"))
				},
			}),
		formatters: formatters,
		pool:       buffer.NewPool(),
	}
}

func (fe *FormatEncoder) Clone() zapcore.Encoder {
	return &FormatEncoder{
		Encoder:    fe.Encoder.Clone(),
		formatters: fe.formatters,
		pool:       fe.pool,
	}
}

func (f *FormatEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 获取一个缓冲区
	line := f.pool.Get()
	for _, f := range f.formatters {
		f.Format(line, entry, fields)
	}
	encodeFields, err := f.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	if line.Len() > 0 && encodeFields.Len() != 1 {
		line.AppendString(" ")
	}
	line.AppendString(encodeFields.String())
	encodeFields.Free()
	return line, nil
}

// StringFormatter 对字符串进行格式化。
type StringFormatter struct{
	Value string
}

// Format 只是将一个字符串写入到io.Writer里，根本没用到传入的entry和fields。
func (s StringFormatter) Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field) {
	fmt.Fprintf(w, "%s", s.Value)
}