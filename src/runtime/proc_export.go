package runtime

import (
	"runtime/internal/atomic"
	"unsafe"
)

var bindps []*bindPG
var bindStats []int64

type (
	bindPG struct {
		gp  *g
		fn  func(argp unsafe.Pointer)
		arg unsafe.Pointer
		pid int
	}
	bindpFuncval struct {
		fn  uintptr
		arg unsafe.Pointer
	}
)

func NewProcBindp(fn func(argp unsafe.Pointer), arg unsafe.Pointer, pid int) {
	funcvalp := *(**funcval)(unsafe.Pointer(&fn))
	gp := getg()
	pc := getcallerpc()
	systemstack(func() {
		newg := newproc1(funcvalp, unsafe.Pointer(&arg), 8, gp, pc)
		newg.isbindp = true
		newg.bindpid = pid
		if atomic.Casint64(&bindStats[pid], 0, -1) {
			bindps[pid] = &bindPG{gp: newg, fn: fn, arg: arg, pid: pid}
			if !atomic.Casint64(&bindStats[pid], -1, 1) {
				throw("set newbindp g to bound pid")
			}
		} else {
			throw("set newbindp g to bound pid")
		}
		allp[pid].bind = true
		startm(allp[pid], true)
	})
}

func GetPid() int32 {
	return getg().m.p.ptr().id
}
