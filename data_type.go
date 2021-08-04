package main

import (
	"fmt"
	//Kalo di JS pake type untum log tipe data, di GO ada reflect.type()
	"reflect"
)

/** Tipe Data di Go-Lang (So far mostly same with other language)
## [ NUMERIK ]
- uint8 (0-255)
- uint16 (0-65535/2^16)
- uint32 (0-2^32)
- uint64 (0-2^64)
- uint (Sama dengan uint32 dan uint64)
- byte (sama dengan uint8)
- int8 (-128 – 127)
- int16 (-32768 – 32767)
- int32 (-2147483648 – 2147483647)
- int64 (-9223372036854775808 – 9223372036854775807)
- int (Sama dengan int 32 dan 64)
- rune (Sama dengan int32) => OR also call typedata char

## [ DECIMAL ]
- float32
- float64

## [ BOOLEAN ]
- bool (true/false)

## [ STRING ]
- string (with double quotes or backtick)

?? [ NIL&Zero Value ] => Kaya Falsefy kalo di JS
1. Nil adalah isi sebuah nilai, jika di JS ada null di Go ada nil
2. Zero value adalah isi falsefy selain nil. Semisal di string adalah string kosong, bool adalah false, numerik non decimal adala 0 dst

## [ Lainnya (Non Primitif tipe data) ]
1. pointer => karena tidak copy by reference tapi copy as value
2. tipe data fungsi
3. array
4. slice
5. map
6. channel
7. interface kosong atau interface {}
8. struct => kaya objek kalo di JS

*/

func main() {
	// ## [ NUMERIK ]
	// coba := -2039039309309498589
	// var coba2 int = 25598989
	// ## [ DECIMAL ]
	// var coba float32 = 7.5
	// ## [ BOOLEAN ]
	// var coba bool = true
	// ## [ STRING ]
	// var coba string = "o"
	// ## [ NON Primitive ]
	// coba := []int{} => Slice
	coba := map[string]int{}
	coba["satu"] = 9
	// coba := [...]int{} => Array
	fmt.Println("Tipe Data =>", reflect.TypeOf(coba), "|| VALUE =>", coba["satu"])

}
