package hrf

import (
	"github.com/batchcorp/plz/countlog/spi"
	"github.com/batchcorp/plz/msgfmt"
)

type ctxFormatter struct {
	fmt msgfmt.Formatter
}

func (formatter *ctxFormatter) Format(space []byte, event *spi.Event) []byte {
	return formatter.fmt.Format(space, spi.GetLogContext(event.Context).Properties)
}
