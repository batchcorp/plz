package main

import (
	"github.com/batchcorp/plz/countlog/output"
	. "github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/countlog/output/printf"
)

func main() {
	EventWriter = output.NewEventWriter(output.EventWriterConfig{
		Format: &printf.Format{
			`[{level}] ` +
				`{timestamp, goTime, 15:04:05} ` +
				`{message} @ {file}:{line}`},
	})
	Info("{userA} called {userB} at {sometime}",
		"userA", "lily",
		"userB", "tom",
		"sometime", "yesterday")
	Info("{userA} called {userB} at {sometime}",
		"userA", "lily",
		"userB", "tom",
		"sometime", "yesterday")
}
