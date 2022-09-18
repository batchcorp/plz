package hrf

import (
	"github.com/batchcorp/plz/countlog/spi"
	"github.com/batchcorp/plz/msgfmt"
)

type propFormatter struct {
	fmt msgfmt.Formatter
}

func (formatter *propFormatter) Format(space []byte, event *spi.Event) []byte {
	return formatter.fmt.Format(space, event.Properties)
}
