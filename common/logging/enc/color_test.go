package enc

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReset(t *testing.T) {
	require.Equal(t, ResetColor(), "\x1b[0m")
}

func TestNormalColors(t *testing.T) {
	require.Equal(t, ColorBlack.Normal(), "\x1b[30m")
	require.Equal(t, ColorRed.Normal(), "\x1b[31m")
	require.Equal(t, ColorGreen.Normal(), "\x1b[32m")
	require.Equal(t, ColorYellow.Normal(), "\x1b[33m")
	require.Equal(t, ColorBlue.Normal(), "\x1b[34m")
	require.Equal(t, ColorMagenta.Normal(), "\x1b[35m")
	require.Equal(t, ColorCyan.Normal(), "\x1b[36m")
	require.Equal(t, ColorWhite.Normal(), "\x1b[37m")
}

func TestBoldColors(t *testing.T) {
	require.Equal(t, ColorBlack.Bold(), "\x1b[30;1m")
	require.Equal(t, ColorRed.Bold(), "\x1b[31;1m")
	require.Equal(t, ColorGreen.Bold(), "\x1b[32;1m")
	require.Equal(t, ColorYellow.Bold(), "\x1b[33;1m")
	require.Equal(t, ColorBlue.Bold(), "\x1b[34;1m")
	require.Equal(t, ColorMagenta.Bold(), "\x1b[35;1m")
	require.Equal(t, ColorCyan.Bold(), "\x1b[36;1m")
	require.Equal(t, ColorWhite.Bold(), "\x1b[37;1m")
}

func TestPrintColor(t *testing.T) {
	w := &bytes.Buffer{}
	fmt.Fprint(w, ResetColor(), "你好", ColorBlue.Normal(), "你也好", ColorMagenta.Normal(), "很有趣", ResetColor())
	t.Log(w.String())
}
