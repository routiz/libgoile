package libgoile_test

import (
	"testing"

	"github.com/routiz/libgoile"
)

func TestEval(t *testing.T) {
	libgoile.ScmInitGuile()

	t.Run("test-eval-string", func(t *testing.T) {
		adderscm := libgoile.ScmEvalString("(lambda (x y) (+ x y))")
		addedscm := libgoile.ScmCall2(adderscm,
			libgoile.ScmFromInt64(10),
			libgoile.ScmFromInt64(20))
		added := libgoile.ScmToInt64(addedscm)
		if added != 30 {
			t.Log("failed to call scheme adder")
			t.Fail()
		}
	})

	t.Run("test-eval-string", func(t *testing.T) {
		scmExp3plus10 := libgoile.ScmEvalString("'(+ 3 10)")
		scm13 := libgoile.ScmEval(scmExp3plus10, nil)
		if libgoile.ScmToInt64(scm13) != 13 {
			t.Fail()
		}
	})
}
