package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"

func ScmThrow(key string, scmlist CSCM) CSCM {
	return CSCM(C.scm_throw(
		ScmFromStringSymbol(key),
		scmlist,
	))
}
