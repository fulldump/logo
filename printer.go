package logo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// defaultPrinter format the output in this way:
// {
//   "time":"<TIMESTAMP>",
//   "type":"LOG",
//   "level":"ERROR",
//   "payload":"Flush to disk failed",
// }
func defaultPrinter(l *Logo, lvl Level, a ...interface{}) {
	e := Entry{
		Time:    time.Now(),
		Type:    "LOG",
		Level:   lvl.String(),
		Payload: strings.TrimRight(fmt.Sprintln(a...), "\n"),
	}

	fmt.Fprint(l.Output, "LOG:")
	json.NewEncoder(l.Output).Encode(e)
}

type Entry struct {
	Time    time.Time `json:"time"`
	Type    string    `json:"type"`
	Level   string    `json:"level"`
	Payload string    `json:"payload"`
}

// Hack format time.Time http://choly.ca/post/go-json-marshalling/
func (this Entry) MarshalJSON() ([]byte, error) {
	type Alias Entry
	return json.Marshal(&struct {
		Time string `json:"time"`
		Alias
	}{
		Time:  this.Time.UTC().Format(time.RFC3339Nano),
		Alias: (Alias)(this),
	})
}
