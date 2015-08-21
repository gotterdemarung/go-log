package main

import (
	"fmt"
	"github.com/gotterdemarung/go-log/log"
)

func main() {
	log.Dispatcher.FromCli()

	l := log.Context.With("test")
	l.Trace("This is trace")
	l.Debug("This is debug")
	l.Info("This is common info")
	l.Warn("This is warning")
	l.Fail(fmt.Errorf("%s", "Some error message"))

	l.Warn("This is warning")
	l.LogWithContext(log.INFO, "Some user :id with name :name, is active", log.Pair{"name", "Foo"}, log.Pair{"id", func() float32 {return 0.4}})
}
