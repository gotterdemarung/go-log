package log

type Level int
const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
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
