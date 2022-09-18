package dump

import (
	"context"
	"encoding/json"
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"reflect"
	"unsafe"
)

var strEncoderInst = jsonfmt.EncoderOf(reflect.TypeOf(""))

type stringHeader struct {
	data unsafe.Pointer
	len  int
}

type stringEncoder struct {
}

func (encoder *stringEncoder) Encode(ctx context.Context, space []byte, ptr unsafe.Pointer) []byte {
	header := (*stringHeader)(ptr)
	space = append(space, `{"data":{"__ptr__":"`...)
	ptrStr := ptrToStr(uintptr(header.data))
	space = append(space, ptrStr...)
	space = append(space, `"},"len":`...)
	space = jsonfmt.WriteInt64(space, int64(header.len))
	space = append(space, `}`...)
	elem := strEncoderInst.Encode(ctx, nil, ptr)
	if uintptr(header.data) != 0 {
		addrMap := ctx.Value(addrMapKey).(map[string]json.RawMessage)
		addrMap[ptrStr] = json.RawMessage(elem)
	}
	return space
}
