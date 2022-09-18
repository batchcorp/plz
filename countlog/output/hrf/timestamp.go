package hrf

import (
	"github.com/batchcorp/plz/countlog/spi"
)

type timestampFormatter struct {
}

func (formatter *timestampFormatter) Format(space []byte, event *spi.Event) []byte {
	return event.Timestamp.AppendFormat(space, "15:04:05.000")
}
