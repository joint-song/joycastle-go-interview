package uniqid

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestPseudoEncrypt(t *testing.T) {
	convey.Convey("Define input value", t, func() {
		from, to := 1000, 1010
		// uniqMap := &sync.Map{}
		uniqMap := make(map[uint64]struct{}, to-from+1)
		for i := from; i <= to; i++ {
			convey.Convey(fmt.Sprint("Generate id by value: ", i), func() {
				r := pseudoEncrypt(uint64(i))
				convey.Convey("Id should not present before", func() {
					convey.So(uniqMap, convey.ShouldNotContainKey, r)
				})
				uniqMap[r] = struct{}{}
				convey.Convey(fmt.Sprintf("It should be reversible between input %d and result %d", i, r), func() {
					convey.So(i, convey.ShouldEqual, pseudoEncrypt(r))
				})
			})
		}
	})
}
