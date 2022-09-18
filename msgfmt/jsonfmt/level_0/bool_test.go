package test

import (
	"testing"
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/test/must"
)

func Test_bool(t *testing.T) {
	t.Run("true", test.Case(func(ctx *countlog.Context) {
		must.Equal("true", jsonfmt.MarshalToString(true))
	}))
}
