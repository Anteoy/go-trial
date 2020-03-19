package main // import "github.com/Anteoy/go-trial"

/////
// Simple example of usage
//
//	 go run examples/simple.go
//
/////

import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/Anteoy/go-trial/mmm"
)

type Coordinate struct {
	x, y int
}

func main() {
	// create a new memory chunk that contains 3 Coordinate structures
	//mc, err := mmm.NewMemChunk(Coordinate{}, 3)
	mc, err := mmm.NewMemChunk(unsafe.Pointer(&Coordinate{}), 3)
	if err != nil {
		log.Fatal(err)
	}
	// print 3
	fmt.Println(mc.NbObjects())

	// write {3,9} at index 0, then print {3,9}
	fmt.Println(mc.Write(0, unsafe.Pointer(&Coordinate{3, 9})))
	// write {17,2} at index 1, then print {17,2}
	fmt.Println(mc.Write(1, unsafe.Pointer(&Coordinate{17, 2})))
	// write {42,42} at index 2, then print {42,42}
	fmt.Println(mc.Write(2, unsafe.Pointer(&Coordinate{42, 42})))

	// print {17,2}
	fmt.Println((*Coordinate)(mc.Read(1).(unsafe.Pointer)))
	fmt.Println(*(*Coordinate)(mc.Read(1).(unsafe.Pointer)))
	fmt.Println(*(*Coordinate)(mc.Read(2).(unsafe.Pointer)))
	// print {42,42}
	fmt.Println(*((*Coordinate)(unsafe.Pointer(mc.Pointer(2)))))
	fmt.Println(*((*Coordinate)(unsafe.Pointer(mc.Pointer(2)))))

	// free memory chunk
	if err := mc.Delete(); err != nil {
		log.Fatal(err)
	}

	//test()

	test3()
}

func test3() {
	fmt.Println("start test3")
	//2019/07/21 11:56:38 unsuppported type: "slice"
	mc, err := mmm.NewMemChunk(unsafe.Pointer(&[]byte{}), 3)
	if err != nil {
		log.Fatal(err)
	}
	// print 3
	fmt.Println(mc.NbObjects())

	a := make([]byte, 10)
	for i := 0; i < len(a); i++ {
		a[i] = byte(i)
	}
	// write {3,9} at index 0, then print {3,9}
	fmt.Println(mc.Write(0, unsafe.Pointer(&a)))

	b := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		b[i] = byte(i + 10)
	}
	// write {17,2} at index 1, then print {17,2}
	fmt.Println(mc.Write(1, unsafe.Pointer(&b)))
	// write {42,42} at index 2, then print {42,42}

	c := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		c[i] = byte(i + 100)
	}
	fmt.Println(mc.Write(2, unsafe.Pointer(&c)))

	// print {17,2}
	d := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		d[i] = byte(i + 1020)
	}
	e := (*[]byte)(mc.Read(1).(unsafe.Pointer))
	fmt.Printf("%+v \n", e)
	// print {42,42}
	fmt.Println(*((*[]byte)(unsafe.Pointer(mc.Pointer(2)))))

	// free memory chunk
	if err := mc.Delete(); err != nil {
		log.Fatal(err)
	}
}

func test() {
	fmt.Println("start test2")
	//2019/07/21 11:56:38 unsuppported type: "slice"
	mc, err := mmm.NewMemChunk([]byte{}, 3)
	if err != nil {
		log.Fatal(err)
	}
	// print 3
	fmt.Println(mc.NbObjects())

	a := make([]byte, 10)
	for i := 0; i < len(a); i++ {
		a[i] = byte(i)
	}
	// write {3,9} at index 0, then print {3,9}
	fmt.Println(mc.Write(0, a))

	b := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		b[i] = byte(i + 10)
	}
	// write {17,2} at index 1, then print {17,2}
	fmt.Println(mc.Write(1, b))
	// write {42,42} at index 2, then print {42,42}

	c := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		c[i] = byte(i + 100)
	}
	fmt.Println(mc.Write(2, c))

	// print {17,2}
	d := make([]byte, 10)
	for i := 0; i < len(b); i++ {
		d[i] = byte(i + 1020)
	}
	d, e := mc.Read(1).([]byte)
	fmt.Printf("%+v , %s\n", d, e)
	// print {42,42}
	fmt.Println(*((*[]byte)(unsafe.Pointer(mc.Pointer(2)))))

	// free memory chunk
	if err := mc.Delete(); err != nil {
		log.Fatal(err)
	}
}
