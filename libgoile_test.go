package libgoile_test

import (
	"testing"
	"unsafe"

	"github.com/routiz/libgoile"
)

func TestLibgoile(t *testing.T) {
	greeterorig := "greeterorig"
	libgoile.ScmWithGuile(libgoile.GoScmWithGuileFuncInfo{
		F: func(greeterany any) unsafe.Pointer {
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
		Args: greeterorig,
	})

	libgoile.ScmInitGuile()
}
