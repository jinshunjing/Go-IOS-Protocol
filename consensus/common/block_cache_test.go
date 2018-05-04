package consensus_common

import (
	"github.com/golang/mock/gomock"
	"github.com/iost-official/prototype/core/mocks"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/iost-official/prototype/core/block"
)

func TestBlockCache(t *testing.T) {
	b0 := block.Block{
		Head: block.BlockHead{
			ParentHash: []byte("nothing"),
		},
		Content: []byte("b0"),
	}

	b1 := block.Block{
		Head: block.BlockHead{
			ParentHash: b0.HeadHash(),
		},
		Content: []byte("b1"),
	}

	b2 := block.Block{
		Head: block.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("b2"),
	}

	b2a := block.Block{
		Head: block.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("fake"),
	}

	b3 := block.Block{
		Head: block.BlockHead{
			ParentHash: b2.HeadHash(),
		},
		Content: []byte("b3"),
	}

	b4 := block.Block{
		Head: block.BlockHead{
			ParentHash: b3.HeadHash(),
		},
	}

	ctl := gomock.NewController(t)

	verifier := func(blk *block.Block, chain block.Chain) bool {
		return true
	}

	base := core_mock.NewMockBlockChain(ctl)
	base.EXPECT().Top().AnyTimes().Return(&b0)

	Convey("Test of Block Cache", t, func() {
		Convey("Add:", func() {
			Convey("normal:", func() {
				bc := NewBlockCache(base, 4)
				err := bc.Add(&b1, verifier)
				So(err, ShouldBeNil)
				So(bc.cachedRoot.depth, ShouldEqual, 1)

			})

			Convey("fork and error", func() {
				bc := NewBlockCache(base, 4)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				So(bc.cachedRoot.depth, ShouldEqual, 2)

				verifier = func(blk *block.Block, chain block.Chain) bool {
					return false
				}
				err := bc.Add(&b3, verifier)
				So(err, ShouldNotBeNil)
			})

			Convey("auto push", func() {
				var ans string
				base.EXPECT().Push(gomock.Any()).AnyTimes().Do(func(block *block.Block) error {
					ans = string(block.Content)
					return nil
				})
				verifier = func(blk *block.Block, chain block.Chain) bool {
					return true
				}
				bc := NewBlockCache(base, 3)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b3, verifier)
				bc.Add(&b4, verifier)
				So(ans, ShouldEqual, "b1")
			})
		})

		Convey("Longest chain", func() {
			Convey("no forked", func() {
				bc := NewBlockCache(base, 10)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				ans := string(bc.LongestChain().Top().Content)
				So(ans, ShouldEqual, "b2")
			})

			Convey("forked", func() {
				var bc BlockCache = NewBlockCache(base, 10)

				bc.Add(&b1, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b2, verifier)
				ans := string(bc.LongestChain().Top().Content)
				So(ans, ShouldEqual, "fake")
				bc.Add(&b3, verifier)
				ans = string(bc.LongestChain().Top().Content)
				So(ans, ShouldEqual, "b3")
			})
		})

		Convey("find blk", func() {
			bc := NewBlockCache(base, 10)
			bc.Add(&b1, verifier)
			bc.Add(&b2a, verifier)
			bc.Add(&b2, verifier)
			ans, err := bc.FindBlockInCache(b2a.HeadHash())
			So(err, ShouldBeNil)
			So(string(ans.Content), ShouldEqual, "fake")

			ans, err = bc.FindBlockInCache(b3.HeadHash())
			So(err, ShouldNotBeNil)

		})

	})

}