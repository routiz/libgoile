package libgoile_test

import (
	"testing"

	"github.com/routiz/libgoile"
)

func TestThrow(t *testing.T) {
	t.Skip("exception make exit code 1. " +
		"we need to find other way to test.")
	libgoile.ScmInitGuile()
	libgoile.ScmThrow(
		"test-exception",
		libgoile.ScmList1(
			libgoile.ScmFromString("err-message"),
		),
	)
}
