package libgoile_test

import (
	"testing"

	"github.com/routiz/libgoile"
)

func TestString(t *testing.T) {
	libgoile.ScmInitGuile()
	t.Run("string?", func(t *testing.T) {
		abcscm := libgoile.ScmEvalString("\"abc\"")
		if !libgoile.ScmIsString(abcscm) {
			t.Log("failed to generate string value")
			t.Fail()
		}
	})
	t.Run("scm_to_from_utf8_stringn", func(t *testing.T) {
		message := "hello, scm"
		scm := libgoile.ScmFromString(message)
		if !libgoile.ScmIsString(scm) {
			t.Log("failed to make message into scm string")
			t.Fail()
		}
		strfromscm := libgoile.ScmToString(scm)
		if strfromscm != message {
			t.Log("failed to make scm message into go string")
			t.Fail()
		}
	})
}
