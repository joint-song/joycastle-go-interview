package uniqid_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"joycastle.mobi/go-interview/pkg/uniqid"

	"github.com/smartystreets/goconvey/convey"
)

func TestUniqID(t *testing.T) {
	var (
		uniqMap     = &sync.Map{}
		ctx, cancel = context.WithCancel(context.Background())
		m           = uniqid.NewManager(ctx)
	)
	convey.Convey("Create manager", t, func() {
		for i := 0; i < 20; i++ {
			convey.Convey(fmt.Sprintf("Generate id in round %d", i), func() {
				id := m.NewID()
				convey.Convey(fmt.Sprintf("Id %d should not present before", id), func() {
					_, ok := uniqMap.Load(id)
					convey.So(ok, convey.ShouldBeFalse)
				})
				uniqMap.Store(id, struct{}{})
			})
		}
	})
	cancel()
}
