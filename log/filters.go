package log

func WithCondition(cond func(*Packet) bool, next Appender) Appender {
	return func (l *Packet) {
		if cond(l) {
			next(l)
		}
	}
}

func WithThreshold(level Type, next Appender) Appender {
	return func (l *Packet) {
		if level.LesserOrEq(l.Level) {
			next(l)
		}
	}
}

