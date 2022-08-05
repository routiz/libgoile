package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"
import (
	"unsafe"
)

// ScmIsString checks if this scm is a string.
func ScmIsString(scm unsafe.Pointer) bool {
	return C.scm_string_p(C.SCM(scm)) == C.SCM_BOOL_T
}

// ScmFromString create a guile string from Golang string
func ScmFromString(s string) unsafe.Pointer {
	slen := len(s)
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	scmstr := C.scm_from_utf8_stringn(cs, C.size_t(slen))
	return unsafe.Pointer(scmstr)
}

// ScmToString converts Guile string into Golang string
func ScmToString(scm unsafe.Pointer) string {
	cstr := C.scm_to_utf8_string(C.SCM(scm))
	return C.GoString(cstr)
}

func ScmFromStringSymbol(s string) unsafe.Pointer {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return unsafe.Pointer(C.scm_from_utf8_symbol(cs))
}

func ScmIsList(x unsafe.Pointer) bool {
	return C.scm_list_p(C.SCM(x)) == C.SCM_BOOL_T
}

func ScmList1(elem1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_list_1(C.SCM(elem1)))
}

func ScmList2(elem1, elem2 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_list_2(C.SCM(elem1), C.SCM(elem2)))
}

func ScmList3(elem1, elem2, elem3 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(
		C.scm_list_3(C.SCM(elem1), C.SCM(elem2), C.SCM(elem3)))
}

func ScmList4(elem1, elem2, elem3, elem4 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_list_4(
		C.SCM(elem1),
		C.SCM(elem2),
		C.SCM(elem3),
		C.SCM(elem4)))
}

func ScmList5(elem1, elem2, elem3, elem4, elem5 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_list_5(
		C.SCM(elem1),
		C.SCM(elem2),
		C.SCM(elem3),
		C.SCM(elem4),
		C.SCM(elem5)))
}
