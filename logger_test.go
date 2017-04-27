package logo

import (
	"fmt"

	. "gopkg.in/check.v1"
)

var filters = []Level{LevelOff, LevelFatal, LevelError, LevelWarning, LevelInfo, LevelDebug, LevelTrace, LevelAll}
var levels = []Level{LevelFatal, LevelError, LevelWarning, LevelInfo, LevelDebug, LevelTrace}

func (w *World) Test_Levels(c *C) {

	w.Logo.Level = LevelAll

	component := "LevelsComponent"
	w.Logo.Component = component

	for _, l := range levels {

		w.Logo.log(l, "this", "is", "an", "error")
		e := w.Output.PopLast()

		c.Assert(e.Type, Equals, "LOG")
		c.Assert(e.Payload, DeepEquals, "this is an error")
	}

}

func (w *World) Test_Filtering(c *C) {

	expected := `
OFF:
FATAL:	FATAL
ERROR:	FATAL	ERROR
WARNING:	FATAL	ERROR	WARNING
INFO:	FATAL	ERROR	WARNING	INFO
DEBUG:	FATAL	ERROR	WARNING	INFO	DEBUG
TRACE:	FATAL	ERROR	WARNING	INFO	DEBUG	TRACE
ALL:	FATAL	ERROR	WARNING	INFO	DEBUG	TRACE
`

	result := "\n"
	for _, f := range filters {

		w.Logo.Level = f

		for _, l := range levels {
			w.Logo.log(l, "this ", "is ", "an ", "error")
		}

		result += fmt.Sprint(f, ":")
		for _, e := range w.Output.lines {
			result += fmt.Sprint("\t", e.Level)
		}
		result += "\n"

		w.Output.Clean()
	}

	c.Assert(expected, Equals, result)

}
