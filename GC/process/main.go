package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func GcFinished() *int {
	p := 1
	runtime.SetFinalizer(&p, func(_ *int) {
		println("gc finished")
	})
	return &p
}

// Allocate
// 分配内存
func Allocate() {
	_ = make([]byte, int((1<<20)*0.25))
}
func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	err := trace.Start(f)
	if err != nil {
		log.Println(err)
	}
	defer trace.Stop()
	GcFinished()
	//当完成GC时停止分配
	for n := 1; n < 50; n++ {
		println("#allocate: ", n)
		Allocate()
	}
	println("terminate")

}
