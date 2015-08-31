package log

type Threshold struct {
	AllowedLevel Level
	Next Appender
}

func (t *Threshold) Deliver(l *LogPacket) {
	if t.AllowedLevel.LesserOrEq(l.Level) {
		t.Next(l)
	}
}