package log

// Logger structure
type Logger struct {
	deliver		Appender
	Tags 		[]string
	Context 	map[string]interface{}
}

// Merges two context maps
func mergeCtxs(a map[string]interface{}, b map[string]interface{}) map[string]interface{} {
	c := map[string]interface{}{}
	if len(a) > 0 {
		for k, v := range a {
			c[k] = v
		}
	}
	if len(b) > 0 {
		for k, v := range b {
			c[k] = v
		}
	}

	return c
}

// Returns new logger with tags attached
func (c *Logger) WithTags(tags ...string) *Logger {
	return &Logger{
		// Copy appender
		deliver: c.deliver,

		// Merge context
		Context: mergeCtxs(c.Context, nil),

		// Merge tags
		Tags: append(c.Tags, tags...),
	}
}

// Returns new logger where current logger context expanded by provided one
func (c *Logger) With(context map[string]interface{}) *Logger {
	return &Logger{
		// Copy appender
		deliver: c.deliver,

		// Merge context
		Context: mergeCtxs(c.Context, context),

		// Merge tags
		Tags: c.Tags,
	}
}

func (l *Logger) Log(level Type, line string) {
	p := PacketByLogger(l, level, line)
	l.deliver(p)
}

func (l *Logger) Trace(line string) {
	l.Log(TRACE, line)
}

func (l *Logger) Debug(line string) {
	l.Log(DEBUG, line)
}

func (l *Logger) In(line string) {
	l.Log(IN, line)
}

func (l *Logger) Out(line string) {
	l.Log(OUT, line)
}

func (l *Logger) Info(line string) {
	l.Log(INFO, line)
}

func (l *Logger) Success(line string) {
	l.Log(SUCCESS, line)
}

func (l *Logger) Warn(line string) {
	l.Log(WARN, line)
}

func (l *Logger) Fail(err error) error {
	l.Log(ERROR, err.Error())

	return err
}

func (l *Logger) Error(err error) error {
	l.Log(ERROR, err.Error())

	return err
}
