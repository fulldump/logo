package logo

var level_to_string = map[Level]string{
	LevelOff:     "OFF",
	LevelFatal:   "FATAL",
	LevelError:   "ERROR",
	LevelWarning: "WARNING",
	LevelInfo:    "INFO",
	LevelDebug:   "DEBUG",
	LevelTrace:   "TRACE",
	LevelAll:     "ALL",
}

func (l Level) String() string {
	s := level_to_string[l]
	return s
}

func LevelString(s string) Level {

	for l, v := range level_to_string {
		if s == v {
			return l
		}
	}

	panic("Invalid level string")
}
