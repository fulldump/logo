package logo

import (
	"encoding/json"
	"time"

	. "gopkg.in/check.v1"
)

func (w *World) Test_Json_Serialization(c *C) {
	tz := time.FixedZone("+1", 3600)
	entry := Entry{
		Time:    time.Date(2016, time.November, 24, 18, 25, 15, 34, tz),
		Type:    "LOG",
		Level:   "ERROR",
		Payload: "some message",
	}

	actual, err := json.Marshal(entry)
	c.Assert(err, IsNil)

	expected := `` +
		`{"time":"2016-11-24T17:25:15.000000034Z","type":"LOG","level":"ERROR","payload":"some message"}`
	c.Assert(string(actual), DeepEquals, expected)
}
