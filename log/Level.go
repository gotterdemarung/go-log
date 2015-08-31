package log

type Level int

const (
	TRACE 	Level = 1
	DEBUG	Level = 2
	INFO	Level = 3
	WARN	Level = 4
	ERROR	Level = 5
)

func (l *Level) ToString() string {
	switch *l {
	case TRACE:
		return "trace"
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	}
	return "unknown"
}

func (l Level) LesserOrEq(other Level) bool {
	return l <= other
}