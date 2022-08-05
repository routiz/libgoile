package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"

func ScmThrow(key string, scmlist C.SCM) C.SCM {
	return C.SCM(C.scm_throw(
		ScmFromStringSymbol(key),
		scmlist,
	))
}
