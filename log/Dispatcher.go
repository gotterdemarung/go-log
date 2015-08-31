package log

import (
	"os"
)

func inSlice(list []string, needle string) bool {
	for _, b := range list {
		if b == needle {
			return true
		}
	}
	return false
}

type AppenderDispatcher struct {
	appenders []Appender
	delivery chan *LogPacket
}

func (d *AppenderDispatcher) FromCli() {
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

	ansiAppender := GetAnsiAppender(
		os.Stdout,
		AppenderOptions{
			Precise: logtime,
			Tags: true,
			Bullets: true,
			Colors: !nocolors,
		},
	)

	thresholder := Threshold{
		AllowedLevel: thresholdLevel,
		Next: ansiAppender,
	}


	if logall {
		d.Register(thresholder.Deliver)
	}
}

func (d *AppenderDispatcher) Register(a Appender) {
	d.appenders = append(d.appenders, a)
}

func (d *AppenderDispatcher) Dispatch(l *LogPacket) {
	d.delivery <- l
}

func (d *AppenderDispatcher) accept() {
	for l := range d.delivery {
		for _, a := range d.appenders {
			a(l)
		}
	}
}

func NewDispatcher() *AppenderDispatcher {
	d := AppenderDispatcher{
		appenders: []Appender{},
		delivery: make(chan *LogPacket),
	}

	go d.accept()

	return &d
}

var Dispatcher = NewDispatcher()
