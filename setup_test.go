package logo

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type World struct {
	Logo   *Logo
	Output *JsonBuffer
}

var _ = Suite(&World{})

func (w *World) SetUpTest(c *C) {

	w.Output = NewJsonBuffer()

	w.Logo = NewLogo()
	w.Logo.Level = LevelAll
	w.Logo.Showline = false
	w.Logo.Component = "MyComponent"
	w.Logo.Output = w.Output

}
