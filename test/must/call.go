package must

import (
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/test/testify/assert"
	"reflect"
	"runtime"
)

type ReturnValue []interface{}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

func Call(f interface{}, argObjs ...interface{}) ReturnValue {
	fVal := reflect.ValueOf(f)
	fType := fVal.Type()
	argVals := make([]reflect.Value, len(argObjs))
	for i := 0; i < len(argObjs); i++ {
		if argObjs[i] == nil {
			argVals[i] = reflect.New(fType.In(i)).Elem()
		} else {
			argVals[i] = reflect.ValueOf(argObjs[i])
		}
	}
	retVals := fVal.Call(argVals)
	retObjs := make([]interface{}, len(retVals))
	for i := 0; i < len(retVals); i++ {
		retObjs[i] = retVals[i].Interface()
	}
	if len(retVals) > 0 && retVals[len(retVals)-1].Type().Implements(errorType) {
		errObj := retVals[len(retVals)-1].Interface()
		err, _ := errObj.(error)
		t := test.CurrentT()
		if assert.NoError(t, err) {
			return ReturnValue(retObjs)
		}
		test.Helper()
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			t.Fatal("check failed")
			return ReturnValue(retObjs)
		}
		t.Fatal(test.ExtractFailedLines(file, line))
	}
	return ReturnValue(retObjs)
}

func (ret ReturnValue) Set(objs ...interface{}) {
	for i, obj := range objs {
		reflect.ValueOf(obj).Elem().Set(reflect.ValueOf(ret[i]))
	}
}
