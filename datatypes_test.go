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
	t.Run("list", func(t *testing.T) {
		lst := libgoile.ScmList3(
			libgoile.ScmFromString("first"),
			libgoile.ScmFromString("second"),
			libgoile.ScmFromString("third"),
		)
		len := libgoile.ScmLength(lst)
		if len != 3 {
			t.Log("expected:", 3, "real:", len)
			t.Fail()
		}
		first := libgoile.ScmToString(libgoile.ScmListRef(lst, 0))
		if first != "first" {
			t.Log("expected:", "first", "real:", first)
			t.Fail()
		}
		second := libgoile.ScmToString(libgoile.ScmListRef(lst, 1))
		if second != "second" {
			t.Log("expected:", "second", "real:", second)
			t.Fail()
		}
		third := libgoile.ScmToString(libgoile.ScmListRef(lst, 2))
		if third != "third" {
			t.Log("expected:", "third", "real:", third)
			t.Fail()
		}
	})
}
