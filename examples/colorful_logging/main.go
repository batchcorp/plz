package main

import (
	"os"
	"time"

	. "github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/countlog/output"
	"github.com/batchcorp/plz/countlog/output/async"
	"github.com/batchcorp/plz/countlog/output/hrf"
)

func main() {
	EventWriter = output.NewEventWriter(output.EventWriterConfig{
		Format: &hrf.Format{ShowTimestamp: true},
		Writer: async.NewAsyncWriter(async.AsyncWriterConfig{
			Writer: os.Stderr,
		}),
	})
	Trace("trace should be used in {scenario}",
		"scenario", "unit test",
		"comment", "we love tracing!")
	Debug("debug should be used in {scenario}",
		"scenario", "integration test, debug weird problem in production")
	Info("info should be used as {scenario}",
		"scenario", "the default production logging level")
	Warn("warn should be used when {scenario}",
		"scenario", "err != nil")
	Error("error should be used when {scenario}",
		"scenario", "err != nil returned to user")
	Fatal("fatal is reserved for panic")
	SetMinLevel(LevelDebug)
	if ShouldLog(LevelTrace) {
		Trace("if ShouldLog(LevelTrace) is necessary")
	}
	Trace("without if, the runtime cost is still minimal")
	time.Sleep(time.Second)
}
