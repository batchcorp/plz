package test

import (
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/test/must"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	t.Run("epoch", test.Case(func(ctx *countlog.Context) {
		must.Equal(`"0001-01-01T00:00:00Z"`, jsonfmt.MarshalToString(time.Time{}))
	}))
}
