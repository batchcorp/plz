package test

import (
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/test/must"
	"testing"
)

func Test_bytes(t *testing.T) {
	t.Run("simple", test.Case(func(ctx *countlog.Context) {
		must.Equal(`"hello"`, jsonfmt.MarshalToString([]byte("hello")))
	}))
	t.Run("unicode", test.Case(func(ctx *countlog.Context) {
		must.Equal(`"\xe4\xb8\xad\xe6\x96\x87"`, jsonfmt.MarshalToString([]byte("中文")))
	}))
	t.Run("unicode and control char", test.Case(func(ctx *countlog.Context) {
		must.Equal(`"\xe4\xb8\xad\n\xe6\x96\x87"`, jsonfmt.MarshalToString([]byte("中\n文")))
	}))
}
