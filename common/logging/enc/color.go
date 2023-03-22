package enc

import "fmt"

type Color uint8

const ColorNone Color = 0

const (
	ColorBlack Color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta // 紫红色
	ColorCyan // 青色
	ColorWhite
)

// Normal 就是只添加颜色效果，不增加粗体效果。
func (c Color) Normal() string {
	return fmt.Sprintf("\x1b[%dm", c)
}

// Bold 不仅添加颜色效果，还添加粗体效果。
func (c Color) Bold() string {
	if c == ColorNone {
		return c.Normal()
	}
	return fmt.Sprintf("\x1b[%d;1m", c)
}

// ResetColor 清除之前添加的颜色和粗体效果。
func ResetColor() string {
	return ColorNone.Normal()
}