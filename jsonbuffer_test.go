package logo

import "encoding/json"

/*
 * NOTICE that it is only a stub for testing, this is not thread safe, not production ready
 */

type JsonBuffer struct {
	lines []*Entry
}

func NewJsonBuffer() *JsonBuffer {
	return &JsonBuffer{
		lines: []*Entry{},
	}
}

func (j *JsonBuffer) Write(p []byte) (int, error) {

	v := &Entry{}
	if err := json.Unmarshal(p, &v); nil == err {
		j.lines = append(j.lines, v)
		return len(p), err
	}

	return 0, nil

}

func (j *JsonBuffer) PopLast() *Entry {
	if len(j.lines) > 0 {
		i := j.lines[0]
		j.lines = j.lines[1:]
		return i
	}

	return nil
}

func (j *JsonBuffer) Clean() {
	for len(j.lines) > 0 {
		j.lines = j.lines[1:]
	}
}
