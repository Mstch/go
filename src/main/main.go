package main

import (
	"runtime"
	"time"
	"unsafe"
)

func fuck(p unsafe.Pointer)  {
	println("fuck",p)
	time.Sleep(3*time.Second)
	println("done")
}

func main()	 {
	ch := make(chan int)
	runtime.NewProcBindp(fuck,unsafe.Pointer(uintptr(1)),0)
	<-ch
}
