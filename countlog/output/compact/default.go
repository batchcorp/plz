package compact

import (
	"github.com/batchcorp/plz/countlog/spi"
	"github.com/batchcorp/plz/msgfmt"
)

type defaultFormatter struct {
	fmt msgfmt.Formatter
}

func (formatter *defaultFormatter) Format(space []byte, event *spi.Event) []byte {
	return formatter.fmt.Format(space, event.Properties)
}
