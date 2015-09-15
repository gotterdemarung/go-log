package log

// Represents log packet receiver
type Appender func(l *Packet)