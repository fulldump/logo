package logo

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (w *World) Test_DirectLogFunctions(c *C) {

	w.Logo.Level = LevelAll

	cases := map[Level]func(a ...interface{}){
		LevelFatal:   w.Logo.Fatal,
		LevelError:   w.Logo.Error,
		LevelWarning: w.Logo.Warning,
		LevelInfo:    w.Logo.Info,
		LevelDebug:   w.Logo.Debug,
		LevelTrace:   w.Logo.Trace,
	}

	for l, f := range cases {
		f("this is my log")
		e := w.Output.PopLast()
		c.Assert(LevelString(e.Level), Equals, l)
	}

}

func (w *World) Test_Showline(c *C) {

	w.Logo.Level = LevelAll

	w.Logo.Showline = true
	w.Logo.Error("this is my log")
	e1 := w.Output.PopLast()

	w.Logo.Showline = false
	w.Logo.Error("this is my log")
	e2 := w.Output.PopLast()

	fmt.Println(len(e1.Payload) > len(e2.Payload), Equals, true)
}
