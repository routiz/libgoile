package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"
import (
	"unsafe"
)

// ScmIsString checks if this scm is a string.
func ScmIsString(scm CSCM) bool {
	return C.scm_string_p(scm) == C.SCM_BOOL_T
}

// ScmFromString create a guile string from Golang string
func ScmFromString(s string) CSCM {
	slen := len(s)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return CSCM(C.scm_from_utf8_stringn(cs, C.size_t(slen)))
}

// ScmToString converts Guile string into Golang string
func ScmToString(scm CSCM) string {
	cstr := C.scm_to_utf8_string(scm)
	return C.GoString(cstr)
}

func ScmFromStringSymbol(s string) CSCM {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return CSCM(C.scm_from_utf8_symbol(cs))
}

func ScmIsList(x CSCM) bool {
	return C.scm_list_p(x) == C.SCM_BOOL_T
}

func ScmList1(elem1 CSCM) CSCM {
	return CSCM(C.scm_list_1(elem1))
}

func ScmList2(elem1, elem2 CSCM) CSCM {
	return CSCM(C.scm_list_2(elem1, elem2))
}

func ScmList3(elem1, elem2, elem3 CSCM) CSCM {
	return CSCM(C.scm_list_3(elem1, elem2, elem3))
}

func ScmList4(elem1, elem2, elem3, elem4 CSCM) CSCM {
	return CSCM(C.scm_list_4(elem1, elem2, elem3, elem4))
}

func ScmList5(elem1, elem2, elem3, elem4, elem5 CSCM) CSCM {
	return CSCM(C.scm_list_5(elem1, elem2, elem3, elem4, elem5))
}
