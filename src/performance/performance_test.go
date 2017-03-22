package performance

import (
	"strings"
	"sync"
	"testing"
)

/* x86-64 win10 i3-6100
 * method            len(str)    ns/op    byte/op   allocs/op
 * Str2ByteUnsafe    1           3.12     0         0
 * Str2ByteNormal    1          14.00     0         0
 *
 * Str2ByteUnsafe    8           3.10     0         0
 * Str2ByteNormal    8          13.70     0         0
 *
 * Str2ByteUnsafe    16           3.12    0         0
 * Str2ByteNormal    16          14.30    0         0
 *
 * Str2ByteUnsafe    32           3.13    0         0
 * Str2ByteNormal    32          14.30    0         0
 *
 * Str2ByteUnsafe    64           3.09    0         0
 * Str2ByteNormal    64          79.60    128       2
 *
 * Str2ByteUnsafe    128          3.09    0         0
 * Str2ByteNormal    128        105.00    256       2
 *
 * Str2ByteUnsafe    256          3.09    0         0
 * Str2ByteNormal    256        162.00    512       2
 *
 * Str2ByteUnsafe    512          3.09    0         0
 * Str2ByteNormal    512        260.00    1024      2
 *
 * Str2ByteUnsafe    1024         3.09    0         0
 * Str2ByteNormal    1024       465.00    2048      2
 *
 * Str2ByteUnsafe    2048         3.09    0         0
 * Str2ByteNormal    2048       890.00    4096      2
 *
 * Str2ByteUnsafe    4096         3.09    0         0
 * Str2ByteNormal    4096       1730.00   16384     2
 *
 * Str2ByteUnsafe    8192         3.09    0         0
 * Str2ByteNormal    8192       3480.00   16384     2
 *
 * Str2ByteUnsafe    16384        3.09    0         0
 * Str2ByteNormal    16384      6386.00   32768     2
 *
 * Str2ByteUnsafe    32768        3.09    0         0
 * Str2ByteNormal    32768     12800.00   65536     2
 */

func testStr2ByteUnsafe(str string) {
	b := str2bytes(str)
	_ = bytes2str(b)
}

func testStr2ByteNormal(str string) {
	b := []byte(str)
	_ = string(b)
}

func BenchmarkStr2ByteUnsafe(b *testing.B) {
	b.StopTimer()
	var str = strings.Repeat("a", 1)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testStr2ByteUnsafe(str)
	}
}

func BenchmarkStr2ByteNormal(b *testing.B) {
	b.StopTimer()
	var str = strings.Repeat("a", 1)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testStr2ByteNormal(str)
	}
}

/* x86-64 win10 i3-6100
 * method   len(object)  capacity    ns/op    byte/op   allocs/op
 * array    4            512 	     233      0         0
 * slice    4            512         565      2048      1

 * array    4            1024 	     482      0         0
 * slice    4            1024        1100     4096      1

 * array    8            512 	     337      0         0
 * slice    8            512         962      4096      1

 * array    8            1024  	     651      0         0
 * slice    8            1024        1830     8192      1
 */

const capacity = 1024

func testArrary() [capacity]int32 {
	var d [capacity]int32

	for i := 0; i < len(d); i++ {
		d[i] = 1
	}
	return d
}

func testSlice() []int32 {
	d := make([]int32, capacity)

	for i := 0; i < len(d); i++ {
		d[i] = 1
	}
	return d
}

func BenchmarkArray(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = testArrary()
	}
}

func BenchmarkSlice(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = testSlice()
	}
}

/* x86-64 win10 i3-6100
 * method   len(object)  capacity    ns/op    byte/op   allocs/op
 * noCap    8            10000 	     305211   137       0
 * withCap  8            10000       306019   4         0
 */

func testMap(m map[int]int) {
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
}

func BenchmarkMapNoCap(b *testing.B) {
	b.StopTimer()
	m := make(map[int]int)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testMap(m)
	}
}

func BenchmarkMapWithCap(b *testing.B) {
	b.StopTimer()
	m := make(map[int]int, 10000)
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testMap(m)
	}
}

/* x86-64 win10 i3-6100
 * method   len(object)  capacity    ns/op    byte/op   allocs/op
 * value    8            10000 	     498658   315604    135
 * ptr      8            10000       701364   392808    10125
 */

func testMapWithValue() {
	m := make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
}

func testMapWithPtr() {
	m := make(map[int]*int, 10000)
	for i := 0; i < 10000; i++ {
		v := i
		m[i] = &v
	}
}

func BenchmarkMapWithValue(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testMapWithValue()
	}
}

func BenchmarkMapWithPtr(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testMapWithPtr()
	}
}

/* x86-64 win10 i3-6100
 * method    ns/op    byte/op   allocs/op
 * call      15       0         0
 * defer     44.8     0         0
 */

var mutex sync.Mutex

func testCall() {
	mutex.Lock()
	mutex.Unlock()
}

func testDeferCall() {
	mutex.Lock()
	defer mutex.Unlock()
}

func BenchmarkCall(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testCall()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testDeferCall()
	}
}

/* x86-64 win10 i3-6100
 * method     input_new_object   ns/op    byte/op   allocs/op
 * call          yes             0.57     0         0
 * interface     yes             18.2     8         1
 *
 * call          no              1.43     0         0
 * interface     no              3.20     0         0
 */

type TestInterface interface {
	Test(int)
}

type Data struct {
	x int
}

func (this *Data) Test(x int) {
	this.x += x
}

func testNomralCall(d *Data) {
	d.Test(100)
}

func testInterfaceCall(t TestInterface) {
	t.Test(100)
}

func BenchmarkNoramlCall(b *testing.B) {
	b.StopTimer()
	//d := &Data{}
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testNomralCall(&Data{})
	}
}

func BenchmarkInterfaceCall(b *testing.B) {
	b.StopTimer()
	//d := &Data{}
	b.ReportAllocs()
	b.SetBytes(2)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		testInterfaceCall(&Data{})
	}
}
