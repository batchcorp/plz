package test

import (
	"github.com/batchcorp/plz/reflect2"
	"testing"
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/test/must"
	"github.com/batchcorp/plz/test"
)

func testOp(f func(api reflect2.API) interface{}) func(t *testing.T) {
	return test.Case(func(ctx *countlog.Context) {
		unsafeResult := f(reflect2.ConfigUnsafe)
		safeResult := f(reflect2.ConfigSafe)
		must.Equal(safeResult, unsafeResult)
	})
}
