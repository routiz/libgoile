package libgoile_test

import (
	"testing"

	"github.com/routiz/libgoile"
)

func TestEval(t *testing.T) {
	libgoile.ScmInitGuile()
	adderscm := libgoile.ScmEvalString("(lambda (x y) (+ x y))")
	addedscm := libgoile.ScmCall2(adderscm,
		libgoile.ScmFromInt64(10),
		libgoile.ScmFromInt64(20))
	added := libgoile.ScmToInt64(addedscm)
	if added != 30 {
		t.Log("failed to call scheme adder")
		t.Fail()
	}
}
