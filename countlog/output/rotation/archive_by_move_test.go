package rotation

import (
	"github.com/batchcorp/plz/clock"
	"github.com/batchcorp/plz/countlog"
	"github.com/batchcorp/plz/test"
	"github.com/batchcorp/plz/test/must"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func Test_ArchiveByMove(t *testing.T) {
	t.Run("", test.Case(func(ctx *countlog.Context) {
		os.Remove("/tmp/test.log")
		ioutil.WriteFile("/tmp/test.log", []byte("hello"), 0644)
		archiver := &ArchiveByMove{
			NamingStrategy: &NameByTime{
				Directory: "/tmp",
				Pattern:   "test-2006.log",
			},
		}
		defer clock.ResetNow()
		clock.Now = func() time.Time {
			return time.Unix(0, 0)
		}
		must.Call(archiver.Archive, "/tmp/test.log")
		content := must.Call(ioutil.ReadFile, "/tmp/test-1970.log")[0]
		must.Equal([]byte("hello"), content)
	}))
}
