package libgoile

// #cgo pkg-config: guile-3.0
//
// #include<stdint.h>
// #include<libguile.h>
//
// void* do_goile_scm_with_guile (void*);
// void* goile_scm_with_guile (void*);
// SCM genGreeterSCM(SCM);
import "C"
import (
	"unsafe"
)

type reqType uint64

type reqID uint64

const (
	reqTypeDataStore = reqType(iota)
	reqTypeDataGet
	reqTypeDataClear
)

type req struct {
	t reqType
	p any
}

type resp struct {
	id reqID
	p  any
}

var (
	reqchann  chan req
	respchann chan resp
)

func init() {
	reqchann = make(chan req)
	respchann = make(chan resp)
}

func runDataStoreWorker() {
	respid := reqID(1)
	storage := make(map[reqID]any)
	for {
		req := <-reqchann
		switch req.t {
		case reqTypeDataStore:
			storage[respid] = req.p
			respchann <- resp{id: respid}
			respid++
		case reqTypeDataGet:
			id, ok := req.p.(reqID)
			if !ok {
				respchann <- resp{}
				continue
			}
			d, ok := storage[id]
			if !ok {
				respchann <- resp{}
				continue
			}
			respchann <- resp{id: id, p: d}
		case reqTypeDataClear:
			id, ok := req.p.(reqID)
			if !ok {
				continue
			}
			delete(storage, id)
		}

	}
}

func init() {
	go runDataStoreWorker()
}

func dataStore(d any) reqID {
	reqchann <- req{t: reqTypeDataStore, p: d}
	rsp := <-respchann
	return rsp.id
}

func dataGet(id reqID) any {
	reqchann <- req{t: reqTypeDataGet, p: id}
	rsp := <-respchann
	return rsp.p
}

func dataClear(id reqID) {
	reqchann <- req{t: reqTypeDataClear, p: id}
}

type goScmWithGuileFuncInfo struct {
	F    func(any) unsafe.Pointer
	Args any
}

//export goScmWithGuileFunc
func goScmWithGuileFunc(ctxid uint64) unsafe.Pointer {
	rsp := dataGet(reqID(ctxid))
	fInfo, ok := rsp.(goScmWithGuileFuncInfo)
	if !ok {
		return nil
	}
	return fInfo.F(fInfo.Args)
}

// ScmWithGuile is Go version of scm_with_guile. It runs a function in
// guile mode and hand the return value over. You will call libgoile
// functions via ScmWithGuile unless using ScmInitGuile.
func ScmWithGuile(f func(any) unsafe.Pointer, args any) unsafe.Pointer {
	reqID := dataStore(goScmWithGuileFuncInfo{F: f, Args: args})
	r := C.goile_scm_with_guile(unsafe.Pointer(&reqID))
	dataClear(reqID)
	return r
}

// ScmInitGuile makes the thread in which it is called into guile
// mode.
func ScmInitGuile() {
	C.scm_init_guile()
}

func ScmEvalString(sexpr string) CSCM {
	csexpr := C.CString(sexpr)
	defer C.free(unsafe.Pointer(csexpr))

	scmsexpr := C.scm_from_utf8_stringn(csexpr, C.ulong(len(sexpr)))
	return CSCM(C.scm_eval_string(scmsexpr))
}

func ScmCDefineGsubr(name string, req, opt, rst int, f unsafe.Pointer) CSCM {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	creq := C.int(req)
	copt := C.int(opt)
	crst := C.int(rst)

	ffnc := C.scm_t_subr(f)

	return CSCM(
		C.scm_c_define_gsubr(cname, creq, copt, crst, ffnc),
	)
}

//export genGreeterSCM
func genGreeterSCM(greeter CSCM) CSCM {
	cgreeter := C.scm_to_utf8_stringn(greeter, nil)
	defer C.free(unsafe.Pointer(cgreeter))

	gogreeter := C.GoString(cgreeter)
	gogreet := "hello, " + gogreeter
	cgreet := C.CString(gogreet)
	defer C.free(unsafe.Pointer(cgreet))

	return CSCM(C.scm_from_utf8_string(cgreet))
}

func ScmCDefineGsubrSample() {
	ScmCDefineGsubr("libgoile-sample", 1, 0, 0, C.genGreeterSCM)
}
