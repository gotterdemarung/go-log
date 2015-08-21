package log

import (
	"fmt"
	"time"
)

type LogPacket struct {
	Time time.Time
	Tags []string
	Level Level
	Message string
	Error error
	Values map[string]interface{}
}

func ContextPacket(c *LogContext, l Level) *LogPacket {
	return &LogPacket{
		Time: time.Now(),
		Tags: c.Tags,
		Level: l,
		Values: make(map[string]interface{}),
	}
}

func (lp *LogPacket) SimpleTime() string {
	return fmt.Sprintf(
		"%02d:%02d:%02d",
		lp.Time.Hour(), lp.Time.Minute(), lp.Time.Second(),
	)
}

func (lp *LogPacket) PreciseString() string {
	return fmt.Sprintf(
		"%04d",
		lp.Time.Nanosecond() / 100000,
	)
}


func (lp *LogPacket) ToString() string {
	return fmt.Sprintf(
		"%02d:%02d:%02d.%06d {%s} %s",
		lp.Time.Hour(), lp.Time.Minute(), lp.Time.Second(), lp.Time.Nanosecond(),
		lp.Level.ToString(),
		lp.Message,
	)
}