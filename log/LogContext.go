package log

type LogContext struct {
	delivery Appender
	Tags []string
}


var Context = LogContext{
	delivery: Dispatcher.Dispatch,
	Tags: []string{},
}

func (c *LogContext) With(tags ...string) LogContext {
	return LogContext{
		// Copy appender
		delivery: c.delivery,

		// Merge tags
		Tags: append(c.Tags, tags...),
	}
}

func (l *LogContext) LogSimple(level Level, line string) {
	p := ContextPacket(l, level)
	p.Message = line
	l.delivery(p)
}

func (l *LogContext) Trace(line string) {
	l.LogSimple(TRACE, line)
}

func (l *LogContext) Debug(line string) {
	l.LogSimple(DEBUG, line)
}

func (l *LogContext) Info(line string) {
	l.LogSimple(INFO, line)
}

func (l *LogContext) Warn(line string) {
	l.LogSimple(WARN, line)
}

func (l *LogContext) Fail(err error) {
	l.LogSimple(ERROR, err.Error())
}

func (l *LogContext) LogWithContext(level Level, line string, pairs ...Pair) {
	p := ContextPacket(l, level)
	p.Message = line + "\n"
	for _, pair := range pairs {
		p.Values[pair.Name] = pair.Value
	}
	l.delivery(p)
}

func (l *LogContext) Tracec(line string, pairs ...Pair) {
	l.LogWithContext(TRACE, line, pairs...)
}

func (l *LogContext) Debugc(line string, pairs ...Pair) {
	l.LogWithContext(DEBUG, line, pairs...)
}

func (l *LogContext) Infoc(line string, pairs ...Pair) {
	l.LogWithContext(INFO, line, pairs...)
}

func (l *LogContext) Warnc(line string, pairs ...Pair) {
	l.LogWithContext(WARN, line, pairs...)
}

func (l *LogContext) Failc(line string, pairs ...Pair) {
	l.LogWithContext(ERROR, line, pairs...)
}