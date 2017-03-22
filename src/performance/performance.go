package performance

import (
	//"fmt"
	//"strings"
	"reflect"
	"unsafe"
)

func str2bytes(s string) []byte {
	//x := (*[2]uintptr)(unsafe.Pointer(&s))
	//h := [3]uintptr{x[0], x[1], x[1]}
	h := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&s)), Len: len(s), Cap: len(s)}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type MyStruct struct {
	A int
	B int
}

var sizeOfMyStruct = int(unsafe.Sizeof(MyStruct{}))

func MyStructToBytes(s *MyStruct) []byte {
	x := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&s)), Len: sizeOfMyStruct, Cap: sizeOfMyStruct}
	/*var x reflect.SliceHeader
	x.Len = sizeOfMyStruct
	x.Cap = sizeOfMyStruct
	x.Data = uintptr(unsafe.Pointer(s))*/
	return *(*[]byte)(unsafe.Pointer(&x))
}

func BytesToMyStruct(b []byte) *MyStruct {
	return (*MyStruct)(unsafe.Pointer(
		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
	))
}
