package logo

type Level int

const (
	LevelOff Level = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace
	LevelAll
)

type Logger interface {
	Log(Level, ...interface{})

	Fatal(...interface{})
	Error(...interface{})
	Warning(...interface{})
	Info(...interface{})
	Debug(...interface{})
	Trace(...interface{})
}
