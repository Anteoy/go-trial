package main

/*
   #include <stdio.h>
   #include <stdlib.h>
   #include <string.h>
   #include <ctype.h>
   struct t {
   char *s;
   };

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var t C.struct_t
	var s = "hello world"
	var ch *C.char
	var tmp *C.char

	// 分配空间, 并判断是否分配成功
	t.s = (*C.char)(C.malloc(C.size_t(100)))
	if t.s == nil {
		//if t.s == (*C.char)(unsafe.Pointer(uintptr(0))) {
		panic("malloc failed!\n")
	}

	// 释放内存
	defer C.free(unsafe.Pointer(t.s))
	// 将go的字符串转为c的字符串，并自动释放
	ch = C.CString(s)
	defer C.free(unsafe.Pointer(ch))

	// 调用C的strncpy函数复制
	C.strncpy(t.s, ch, C.size_t(len(s)))
	// C的指针操作
	for i := C.size_t(0); i < C.strlen(t.s); i++ {
		tmp = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(t.s)) + uintptr(i)))
		*tmp = C.char(C.toupper(C.int(*tmp)))
	}

	fmt.Printf("%s\n", C.GoString(t.s))
	fmt.Printf("sizeof struct t is %v\n", unsafe.Sizeof(t))
}
