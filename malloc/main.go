package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"github.com/Anteoy/go-trial/malloc/cmx"
)

const length int = 10

func main() {
	for {
		//item.Body = make([]byte, length)
		mm := cmx.Malloc(length)
		data := (*[1 << 30]byte)(mm)[:length]
		(*reflect.SliceHeader)(unsafe.Pointer(&data)).Cap = length

		for i := 0; i < length; i++ {
			data[i] = byte(i)
		}
		time.Sleep(time.Second * 1)
		fmt.Println("DATA: ", data)
		//fmt.Println(1111)
		cmx.Free(mm)
	}

}
