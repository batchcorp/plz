package countlog

import (
	"os"
	"github.com/batchcorp/plz/countlog/output"
	"github.com/batchcorp/plz/countlog/stats"
	"github.com/batchcorp/plz/countlog/output/hrf"
	"github.com/batchcorp/plz/countlog/spi"
)

var EventWriter spi.EventSink = output.NewEventWriter(output.EventWriterConfig{
	Format: &hrf.Format{},
	Writer: os.Stdout,
})

var EventAggregator spi.EventSink = stats.NewEventAggregator(stats.EventAggregatorConfig{
	Collector: nil, // set Collector to enable stats
})
