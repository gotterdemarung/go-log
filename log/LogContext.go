package log

type LogContext struct {
	delivery Appender
	Tags []string
	Context map[string]interface{}
}

var Context = LogContext{
	delivery: Dispatcher.Dispatch,
	Tags: []string{},
	Context: map[string]interface{}{},
}

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

func (c *LogContext) WithTags(tags ...string) *LogContext {
	return &LogContext{
		// Copy appender
		delivery: c.delivery,

		// Merge context
		Context: mergeCtxs(c.Context, nil),

		// Merge tags
		Tags: append(c.Tags, tags...),
	}
}

func (c *LogContext) WithContext(context map[string]interface{}) *LogContext {
	return &LogContext{
		// Copy appender
		delivery: c.delivery,

		// Merge context
		Context: mergeCtxs(c.Context, context),

		// Merge tags
		Tags: c.Tags,
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

func (l *LogContext) Fail(err error) error {
	l.LogSimple(ERROR, err.Error())

	return err
}

func (l *LogContext) LogWithContext(level Level, line string, context map[string]interface{}) {
	p := ContextPacket(l, level)
	p.Message = line
	p.Values = mergeCtxs(l.Context, context)
	l.delivery(p)
}

func (l *LogContext) Tracec(line string, context map[string]interface{}) {
	l.LogWithContext(TRACE, line, context)
}

func (l *LogContext) Debugc(line string, context map[string]interface{}) {
	l.LogWithContext(DEBUG, line, context)
}

func (l *LogContext) Infoc(line string, context map[string]interface{}) {
	l.LogWithContext(INFO, line, context)
}

func (l *LogContext) Warnc(line string, context map[string]interface{}) {
	l.LogWithContext(WARN, line, context)
}

func (l *LogContext) Failc(line string, context map[string]interface{}) {
	l.LogWithContext(ERROR, line, context)
}
