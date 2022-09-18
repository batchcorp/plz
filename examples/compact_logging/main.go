package main

import (
	. "github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/countlog/output"
	"github.com/batchcorp/plz/countlog/output/async"
	"github.com/batchcorp/plz/countlog/output/json"
	"github.com/batchcorp/plz/countlog/output/lumberjack"
)

func main() {
	writer := async.NewAsyncWriter(async.AsyncWriterConfig{
		Writer: &lumberjack.Logger{
			Filename: "/tmp/test.log.json",
		},
	})
	defer writer.Close()
	EventWriter = output.NewEventWriter(output.EventWriterConfig{
		Format: &json.JsonFormat{},
		Writer: writer,
	})
	for i := 0; i < 10; i++ {
		Info("game score calculated",
			"playerId", 1328+i,
			"scores", []int{1, 2, 7 + i})
	}
}
