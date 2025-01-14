package test

import (
	"testing"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/test/must"
	"github.com/batchcorp/plz/dump"
)

func Test_level2(t *testing.T) {
	t.Run("slice of string", test.Case(func(ctx *countlog.Context) {
		must.JsonEqual(`{
		"__root__": {
			"type": "[]string",
			"data": {
				"__ptr__": "{ptr1}"
			}
		},
		"{ptr1}": {
			"data": {
				"__ptr__": "{ptr2}"
			},
			"len": 2,
			"cap": 2
		},
		"{ptr2}": [
			{
				"data": {
					"__ptr__": "{ptr3}"
				},
				"len": 5
			},
			{
				"data": {
					"__ptr__": "{ptr4}"
				},
				"len": 5
			}
		],
		"{ptr3}": "hello",
		"{ptr4}": "world"}`, dump.Var{[]string{
			"hello",
			"world",
		}}.String())
	}))
}