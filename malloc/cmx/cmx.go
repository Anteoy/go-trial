package cmx

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Malloc(size int) unsafe.Pointer {
	return (unsafe.Pointer)(C.malloc(C.ulong(size)))
}

func Free(ptr unsafe.Pointer) {
	C.free(unsafe.Pointer(ptr))
}
