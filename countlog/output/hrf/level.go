package hrf

import (
	"github.com/batchcorp/plz/countlog/spi"
)

type levelFormatter struct {
}

func (formatter *levelFormatter) Format(space []byte, event *spi.Event) []byte {
	return append(space, spi.ColoredLevelName(event.Level)...)
}
