package log

import (
	"fmt"
	"time"
)

// Log packet structure
type Packet struct {
	Time time.Time
	Tags []string
	Level Type
	Message string
	Error error
	Values map[string]interface{}
}

// Returns common packet
func PacketByLogger(c *Logger, l Type, message string) *Packet {
	return &Packet{
		Time: time.Now(),
		Tags: c.Tags,
		Level: l,
		Message: message,
		Values: c.Context,
	}
}

// Returns true if packet contains requested tag
func (p *Packet) HasTag(name string) bool {
	for _, t := range p.Tags {
		if t == name {
			return true
		}
	}

	return false
}

// Returns packet time in hh:mm:ss format
func (lp *Packet) SimpleTime() string {
	return fmt.Sprintf(
		"%02d:%02d:%02d",
		lp.Time.Hour(), lp.Time.Minute(), lp.Time.Second(),
	)
}

// Returns packet milliseconds only (4 digits after point)
func (lp *Packet) PreciseString() string {
	return fmt.Sprintf(
		"%04d",
		lp.Time.Nanosecond() / 100000,
	)
}

// String representation of packet, for debug purposes only
// Appenders must perform this operation manually
func (lp *Packet) ToString() string {
	return fmt.Sprintf(
		"%02d:%02d:%02d.%06d {%s} %s",
		lp.Time.Hour(), lp.Time.Minute(), lp.Time.Second(), lp.Time.Nanosecond(),
		lp.Level.ToString(),
		lp.Message,
	)
}