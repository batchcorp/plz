package test

import (
	"github.com/batchcorp/plz/msgfmt/jsonfmt"
	"github.com/batchcorp/plz/reflect2"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_slice(t *testing.T) {
	should := require.New(t)
	encoder := jsonfmt.EncoderOf(reflect2.TypeOf([]int(nil)))
	should.Equal("[1,2,3]", string(encoder.Encode(nil, nil, reflect2.PtrOf([]int{
		1, 2, 3,
	}))))
}
