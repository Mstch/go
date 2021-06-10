package main

import (
	"runtime"
	"time"
	"unsafe"
)

func fuck(p unsafe.Pointer)  {
	println(p)
	time.Sleep(3*time.Second)
	println("done")
}

func main()	 {
	ch := make(chan int)
	println("main",&ch)
	runtime.NewProcBindp(fuck,unsafe.Pointer(&ch),0)
	runtime.NewProcBindp(fuck,unsafe.Pointer(&ch),1)
	<-ch
}
