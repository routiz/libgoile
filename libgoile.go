package libgoile

// #cgo pkg-config: guile-3.0
//
// #include<stdint.h>
// #include<libguile.h>
//
// void* do_goile_scm_with_guile (void*);
// void* goile_scm_with_guile (void*);
import "C"
import (
	"unsafe"
)

type reqType uint64

type reqID uint64

const (
	reqTypeDataStore = reqType(iota)
	reqTypeDataGet
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
	respid := reqID(1)
	storage := make(map[reqID]any)
	go func() {
		for {
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
				}

			}
		}
	}()
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

type goScmWithGuileFuncInfo struct {
	F    func(any) unsafe.Pointer
	Args any
}

//export GoScmWithGuileFunc
func GoScmWithGuileFunc(ctxid uint64) unsafe.Pointer {
	rsp := dataGet(reqID(ctxid))
	fInfo, ok := rsp.(goScmWithGuileFuncInfo)
	if !ok {
		return nil
	}
	return fInfo.F(fInfo.Args)
}

func ScmWithGuile(f func(any) unsafe.Pointer, args any) unsafe.Pointer {
	reqID := dataStore(goScmWithGuileFuncInfo{F: f, Args: args})
	return C.goile_scm_with_guile(unsafe.Pointer(&reqID))
}

func ScmInitGuile() {
	C.scm_init_guile()
}

func ScmEvalString(sexpr string) Scm {
	csexpr := C.CString(sexpr)
	defer C.free(unsafe.Pointer(csexpr))

	scmsexpr := C.scm_from_utf8_stringn(csexpr, C.ulong(len(sexpr)))

	scmp := C.scm_eval_string(scmsexpr)
	return Scm{p: unsafe.Pointer(scmp)}
}
