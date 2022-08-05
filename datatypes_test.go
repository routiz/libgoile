package libgoile_test

import (
	"testing"

	"github.com/routiz/libgoile"
)

func TestDatatype(t *testing.T) {
	t.Run("integer", func(t *testing.T) {
		libgoile.ScmInitGuile()

		scm := libgoile.ScmFromInt64(int64(12345))
		if !libgoile.ScmIsNumber(scm) {
			t.Log("12345 is number")
			t.Fail()
		}
		if !libgoile.ScmIsInteger(scm) {
			t.Log("12345 is integer")
			t.Fail()
		}
		if !libgoile.ScmIsExactInteger(scm) {
			t.Log("12345 is exact integer")
			t.Fail()
		}
		intfromscm := libgoile.ScmToInt64(scm)
		if intfromscm != 12345 {
			t.Log("expected:", 12345, "real:", intfromscm)
		}

		scm = libgoile.ScmFromFloat64(float64(5.0))
		if !libgoile.ScmIsNumber(scm) {
			t.Log("5.0 is number")
			t.Fail()
		}
		if !libgoile.ScmIsInteger(scm) {
			t.Log("5.0 is integer")
			t.Fail()
		}
		if libgoile.ScmIsExactInteger(scm) {
			t.Log("5.0 is not exact integer")
			t.Fail()
		}
		ffromscm := libgoile.ScmToFloat64(scm)
		if ffromscm != 5.0 {
			t.Log("expected:", 5.0, "real:", ffromscm)
		}

		scm = libgoile.ScmFromFloat64(float64(5.1))
		if !libgoile.ScmIsNumber(scm) {
			t.Log("5.1 is number")
			t.Fail()
		}
		if libgoile.ScmIsInteger(scm) {
			t.Log("5.1 is not integer")
			t.Fail()
		}
		if libgoile.ScmIsExactInteger(scm) {
			t.Log("5.1 is not exact integer")
			t.Fail()
		}
		ffromscm = libgoile.ScmToFloat64(scm)
		if ffromscm != 5.1 {
			t.Log("expected:", 5.1, "real:", ffromscm)
		}

		scm = libgoile.ScmFromString("12345")
		if libgoile.ScmIsNumber(scm) {
			t.Log("\"12345\" is not number")
			t.Fail()
		}

	})
}
