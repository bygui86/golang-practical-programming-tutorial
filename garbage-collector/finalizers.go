package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

type Foo struct {
	a int
}

func main() {
	debug.SetGCPercent(-1)

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	for i := 0; i < 1000000; i++ {
		f := NewFoo(i)
		_ = fmt.Sprintf("%d", f.a)
	}

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	/* INFO
		The finalizers are first removed and the memory will be released on the next cycle.
		We can clearly see that the second run will actually clean the data. The finalizers could therefore slightly affect the performance and memory usage.
	*/

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)
}

//go:noinline
func NewFoo(i int) *Foo {
	f := &Foo{a: rand.Intn(50)}
	runtime.SetFinalizer(f, func(f *Foo) {
		_ = fmt.Sprintf("foo " + strconv.Itoa(i) + " has been garbage collected")
	})

	return f
}
