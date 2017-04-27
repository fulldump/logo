package logo

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

type Logo struct {
	Level     Level
	Showline  bool
	Component string
	Output    io.Writer
	Print     func(*Logo, Level, ...interface{})
}

func NewLogo() *Logo {
	return &Logo{
		Output:   os.Stdout,
		Showline: true,
		Print:    defaultPrinter,
	}
}

func (l *Logo) log(level Level, a ...interface{}) {

	if l.Showline && level != LevelInfo {
		if _, file, line, ok := runtime.Caller(3); ok {
			a = append(a, fmt.Sprintf(" (%s:%d)", file, line))
		}
	}

	if l.Level < level {
		return
	}

	l.Print(l, level, a...)
}

func (l *Logo) Fatal(a ...interface{}) {
	l.log(LevelFatal, a...)
}

func (l *Logo) Error(a ...interface{}) {
	l.log(LevelError, a...)
}

func (l *Logo) Warning(a ...interface{}) {
	l.log(LevelWarning, a...)
}
func (l *Logo) Info(a ...interface{}) {
	l.log(LevelInfo, a...)
}

func (l *Logo) Debug(a ...interface{}) {
	l.log(LevelDebug, a...)
}

func (l *Logo) Trace(a ...interface{}) {
	l.log(LevelTrace, a...)
}

var Default = NewLogo()

var Fatal = Default.Fatal
var Error = Default.Error
var Warning = Default.Warning
var Info = Default.Info
var Debug = Default.Debug
var Trace = Default.Trace
