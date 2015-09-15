package log

// Represents log entry type
type Type int

const (
	TRACE    Type = 10
	DEBUG    Type = 20
	IN		 Type = 21
	OUT		 Type = 22
	INFO     Type = 30
	SUCCESS  Type = 31
	WARN     Type = 40
	ERROR    Type = 50
)

// Returns string representation of type
func (l *Type) ToString() string {
	switch *l {
	case TRACE:
		return "trace"
	case DEBUG:
		return "debug"
	case IN:
		return "in"
	case OUT:
		return "out"
	case INFO:
		return "info"
	case SUCCESS:
		return "success"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	}
	return "unknown"
}

// Returns true if current entry type has priority lesser on equals
// to provided entry type
func (l Type) LesserOrEq(other Type) bool {
	return l <= other
}

// Returns true if current entry type has priority lesser than provided
func (l Type) LesserThan(other Type) bool {
	return l < other
}