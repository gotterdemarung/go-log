package log

import "os"

// Root dispatcher
var Dispatcher *AsyncDispatcher

// Root context
var Context *Logger

// init function
func init() {
	Dispatcher = NewAsyncDispatcher()
	Context = &Logger{
		deliver: Dispatcher.Dispatch,
		Tags: []string{},
		Context: map[string]interface{}{},
	}
}

func Autoconfig() {
	args := os.Args[1:]

	logall := inSlice(args, "--logall")
	logtime := inSlice(args, "--logtime")
	nocolors := inSlice(args, "--nocolor") || inSlice(args, "--no-color") || inSlice(args, "--no-ansi")

	thresholdLevel := INFO
	if inSlice(args, "-vv") || inSlice(args, "-vvv") {
		thresholdLevel = TRACE
	} else if inSlice(args, "-v") {
		thresholdLevel = DEBUG
	}

	stdoutAppender := GetAnsiAppender(
		os.Stdout,
		AnsiAppenderOptions{
			Precise: logtime,
			Tags: true,
			Bullets: true,
			Colors: !nocolors,
		},
	)

	stderrAppender := GetAnsiAppender(
		os.Stderr,
		AnsiAppenderOptions{
			Precise: logtime,
			Tags: true,
			Bullets: true,
			Colors: !nocolors,
		},
	)

	Dispatcher.Register(WithThreshold(WARN, stderrAppender))
	if logall {
		Dispatcher.Register(WithCondition(func(l *Packet) bool {
			return thresholdLevel.LesserOrEq(l.Level) && l.Level.LesserThan(WARN)
		}, stdoutAppender))
	} else {

	}
}

func inSlice(list []string, needle string) bool {
	for _, b := range list {
		if b == needle {
			return true
		}
	}
	return false
}

