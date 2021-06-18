package main

import (
	"runtime"
	"time"
	"unsafe"
)

func fuck(p unsafe.Pointer) {
	println("fuck", unsafe.Pointer(&p), uintptr(p))
	time.Sleep(3 * time.Second)
	println("done")
	println("")
}

func main() {
	ch := make(chan int)
	runtime.NewProcBindp(fuck,unsafe.Pointer(unsafe.Pointer(uintptr(1))),0)
	<-ch
}
