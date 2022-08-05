package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"
import "unsafe"

func ScmThrow(key string, scmlist unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_throw(
		C.SCM(ScmFromStringSymbol(key)),
		C.SCM(scmlist),
	))
}
