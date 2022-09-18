package msgfmt

import (
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"github.com/batchcorp/plz/reflect2"
)

type jsonFormatter struct {
	idx     int
	encoder jsonfmt.Encoder
}

func (formatter *jsonFormatter) Format(space []byte, kv []interface{}) []byte {
	return formatter.encoder.Encode(nil, space, reflect2.PtrOf(kv[formatter.idx]))
}
