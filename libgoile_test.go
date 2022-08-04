package libgoile_test

import (
	"testing"
	"unsafe"

	"github.com/routiz/libgoile"
)

func TestScmWithGuile(t *testing.T) {
	greeterorig := "greeterorig"
	libgoile.ScmWithGuile(
		func(greeterany any) unsafe.Pointer {
			greeter, ok := greeterany.(string)
			if !ok {
				t.Logf("Failed to get args: %T\n", greeter)
				t.Fail()
				return nil
			}
			if greeter != greeterorig {
				t.Log("expected:", greeterorig, "real:", greeter)
				t.Fail()
			}
			return nil
		},
		greeterorig)

	libgoile.ScmWithGuile(
		func(greeterany any) unsafe.Pointer {
			libgoile.ScmEvalString("(+ 2 10)")
			// TODO When type Scm struct{} is done, check
			// the result.
			return nil
		},
		nil)
}

func TestScmInitGuile(t *testing.T) {
	libgoile.ScmInitGuile()
}

func TestDefineSubr(t *testing.T) {
	libgoile.ScmWithGuile(
		func(any) unsafe.Pointer {
			libgoile.ScmCDefineGsubrSample()
			libgoile.ScmEvalString("(libgoile-sample \"libgoile\")")
			return nil
		}, nil,
	)
}
