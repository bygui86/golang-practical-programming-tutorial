package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	simpleExamples()
	PrintMemUsage()

	structExamples()
	PrintMemUsage()

	listExamples()
	PrintMemUsage()
}

func simpleExamples() {
	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a)) // 8
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b)) // 8
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c)) // 16
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d)) // 24
	fmt.Printf("d: %T, %d\n", d, DeepSize(d))      // 24
}

func structExamples() {
	tuple := Tuple{
		open:     42.42,
		high:     42.42,
		low:      42.42,
		close:    42.42,
		volume:   42.42,
		openedAt: time.Now(),
		closedAt: time.Now(),
	}
	fmt.Printf("tuple unsafe sizeof: %d\n", unsafe.Sizeof(tuple))        // 88
	fmt.Printf("tuple reflect size: %d\n", reflect.TypeOf(tuple).Size()) // 88
	fmt.Printf("tuple deepsize: %d\n", DeepSize(tuple))                  // 88

	now := time.Now()
	tuplePointer := &TuplePointer{
		open:     42.42,
		high:     42.42,
		low:      42.42,
		close:    42.42,
		volume:   42.42,
		openedAt: &now,
		closedAt: &now,
	}
	fmt.Printf("tuple-pointer unsafe sizeof: %d\n", unsafe.Sizeof(*tuplePointer))        // 56
	fmt.Printf("tuple-pointer reflect size: %d\n", reflect.TypeOf(*tuplePointer).Size()) // 56
	fmt.Printf("tuple-pointer deepsize: %d\n", DeepSize(*tuplePointer))                  // 216
}

func listExamples() {
	list := make([]Tuple, 10)
	for i := 0; i < 10; i++ {
		list[i] = Tuple{
			open:     42.42,
			high:     42.42,
			low:      42.42,
			close:    42.42,
			volume:   42.42,
			openedAt: time.Now(),
			closedAt: time.Now(),
		}
	}
	fmt.Printf("list of %d tuples unsafe sizeof: %d\n", len(list), unsafe.Sizeof(list))        // 24
	fmt.Printf("list of %d tuples reflect size: %d\n", len(list), reflect.TypeOf(list).Size()) // 24
	fmt.Printf("list of %d tuples deepsize: %d\n", len(list), DeepSize(list))                  // 904 = 88 (Tuple) * 10 + 24 (slice)

	listPointer := make([]*Tuple, 10)
	for i := 0; i < 10; i++ {
		listPointer[i] = &Tuple{
			open:     42.42,
			high:     42.42,
			low:      42.42,
			close:    42.42,
			volume:   42.42,
			openedAt: time.Now(),
			closedAt: time.Now(),
		}
	}
	fmt.Printf("list of %d pointers to tuples unsafe sizeof: %d\n", len(listPointer), unsafe.Sizeof(listPointer))        // 24
	fmt.Printf("list of %d pointers to tuples reflect size: %d\n", len(listPointer), reflect.TypeOf(listPointer).Size()) // 24
	fmt.Printf("list of %d pointers to tuples deepsize: %d\n", len(listPointer), DeepSize(listPointer))                  // 984 = 88 (Tuple) * 10 + 24 (slice) + 8 (pointer) * 10

	pointerList := make([]*TuplePointer, 10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		pointerList[i] = &TuplePointer{
			open:     42.42,
			high:     42.42,
			low:      42.42,
			close:    42.42,
			volume:   42.42,
			openedAt: &now,
			closedAt: &now,
		}
	}

	// 1 => 248 = 216 + 24 + 8
	// 2 => 312 = (136 * 2) + 24 + (8 * 2)
	// 10 => 824 = (72 * 10) + 24 + (8 * 10)
	// 20 => 1464 = (64 * 20) + 24 + (8 * 20)
	// 100 => 6584 = (57.6 * 100) + 24 + (8 * 100)
	// 1000 => 64184 = (56.16 * 1000) + 24 + (8 * 1000)
	fmt.Printf("list of %d tuplePointers unsafe sizeof: %d\n", len(pointerList), unsafe.Sizeof(pointerList))        // 24
	fmt.Printf("list of %d tuplePointers reflect size: %d\n", len(pointerList), reflect.TypeOf(pointerList).Size()) // 24
	fmt.Printf("list of %d tuplePointers deepsize: %d\n", len(pointerList), DeepSize(pointerList))                  // 824 = (72 * 10) + 24 + (8 * 10)
}

// STRUCTS

type Tuple struct {
	open     float64
	high     float64
	low      float64
	close    float64
	volume   float64
	openedAt time.Time
	closedAt time.Time
}

type TuplePointer struct {
	open     float64
	high     float64
	low      float64
	close    float64
	volume   float64
	openedAt *time.Time
	closedAt *time.Time
}
