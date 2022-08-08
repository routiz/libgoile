package libgoile

import "unsafe"

// #cgo pkg-config: guile-3.0
//
// #include<libguile.h>
import "C"

func ScmEval(exp, env unsafe.Pointer) unsafe.Pointer {
	if env == nil {
		env = unsafe.Pointer(C.scm_interaction_environment())
	}
	return unsafe.Pointer(C.scm_eval(C.SCM(exp), C.SCM(env)))
}

func ScmEvalString(sexpr string) unsafe.Pointer {
	csexpr := C.CString(sexpr)
	defer C.free(unsafe.Pointer(csexpr))

	scmsexpr := C.scm_from_utf8_stringn(csexpr, C.ulong(len(sexpr)))
	return unsafe.Pointer(C.scm_eval_string(scmsexpr))
}

func ScmCall0(proc unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_0(C.SCM(proc)))
}

func ScmCall1(proc, arg1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_1(C.SCM(proc), C.SCM(arg1)))
}

func ScmCall2(proc, arg1, arg2 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_2(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2)))
}

func ScmCall3(proc, arg1, arg2, arg3 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_3(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3)))
}
func ScmCall4(proc, arg1, arg2, arg3, arg4 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_4(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4)))
}
func ScmCall5(proc, arg1, arg2, arg3, arg4, arg5 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_5(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4),
		C.SCM(arg5)))
}
func ScmCall6(proc, arg1, arg2, arg3, arg4, arg5, arg6 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_6(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4),
		C.SCM(arg5), C.SCM(arg6)))
}
func ScmCall7(proc, arg1, arg2, arg3, arg4, arg5, arg6, arg7 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_7(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4),
		C.SCM(arg5), C.SCM(arg6), C.SCM(arg7)))
}
func ScmCall8(proc, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_8(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4),
		C.SCM(arg5), C.SCM(arg6), C.SCM(arg7), C.SCM(arg8)))
}
func ScmCall9(proc, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.scm_call_9(C.SCM(proc),
		C.SCM(arg1), C.SCM(arg2), C.SCM(arg3), C.SCM(arg4),
		C.SCM(arg5), C.SCM(arg6), C.SCM(arg7), C.SCM(arg8), C.SCM(arg9)))
}
