package read_test

import (
	"testing"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/countlog"
	"unicode"
	"github.com/batchcorp/plz/parse/read"
	"github.com/batchcorp/plz/parse"
	"github.com/batchcorp/plz/test/must"
)

func TestUnicodeRanges(t *testing.T) {
	t.Run("complex range", test.Case(func(ctx *countlog.Context) {
		src := parse.NewSourceString("ab中文c,")
		id := read.UnicodeRanges(src, nil, nil, []*unicode.RangeTable{
			unicode.Pattern_Syntax,
			unicode.Pattern_White_Space,
		})
		must.Equal("ab中文c", string(id))
	}))
}
