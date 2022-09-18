package test

import (
	"testing"
	. "github.com/batchcorp/plz/countlog"
	. "github.com/batchcorp/plz/test"
	. "github.com/batchcorp/plz/test/must"
)

func Test(t *testing.T) {
	t.Run("1 != 2", Case(func(ctx *Context) {
		Assert(1 == 2)
	}))
}
