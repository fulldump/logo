package logo

import . "gopkg.in/check.v1"

func (w *World) Test_LevelString_OK(c *C) {

	l := LevelString("INFO")

	c.Assert(l, Equals, LevelInfo)

}

// Invalid Strings should PANIC !
func (w *World) Test_LevelString_Invalid(c *C) {

	c.Assert(func() {
		LevelString("INVALID STRING LOG LEVEL")
	}, PanicMatches, "Invalid level string")
}
