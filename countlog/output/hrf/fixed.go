package hrf

import "github.com/batchcorp/plz/countlog/spi"

type fixedFormatter string

func (formatter fixedFormatter) Format(space []byte, event *spi.Event) []byte {
	return append(space, formatter...)
}