package main

import (
	"fmt"
	"unsafe"
)

// Alignment is about placing fields on address alignment boundaries
// for more efficient reads and writes to memory.

// 对齐是关于将字段放置在地址对齐边界上，以便更有效地读取和写入内存。

// Sample program to show how struct types align on boundaries.
// 显示结构类型如何在边界上对齐的示例程序。

// No byte padding.
type nbp struct {
	a bool // 	1 byte				sizeof 1
	b bool // 	1 byte				sizeof 2
	c bool // 	1 byte				sizeof 3 - Aligned on 1 byte
}

// Single byte padding.
type sbp struct {
	a bool //	1 byte				sizeof 1
	//			1 byte padding		sizeof 2
	b int16 // 	2 bytes				sizeof 4 - Aligned on 2 bytes
}

// Three byte padding.
type tbp struct {
	a bool //	1 byte				size 1
	//			3 bytes padding		size 4
	b int32 //	4 bytes				size 8 - Aligned on 2 bytes
}

// Seven byte padding.
type svnbp struct {
	a bool //	1 byte				size 1
	//			7 bytes padding		size 8
	b int64 //	8 bytes				size 16 - Aligned on 8 bytes
}

// Eight byte padding on 64bit Arch. Word size is 8 bytes.
type ebp64 struct {
	a string //	16 bytes			size 16
	b int32  //	 4 bytes			size 20
	//  		 4 bytes padding	size 24
	c string //	16 bytes			size 40
	d int32  //	 4 bytes			size 44
	//  		 4 bytes padding	size 48 - Aligned on 8 bytes
}

// No padding on 32bit Arch. Word size is 4 bytes.
// To see this build as 32 bit: GOARCH=386 go build
type ebp32 struct {
	a string //	 8 bytes			size  8
	b int32  //	 4 bytes			size 12
	c string //	 8 bytes			size 20
	d int32  //	 4 bytes			size 24 - Aligned on 4 bytes
}

// No padding.
type np struct {
	a string // 16 bytes			size 16
	b string // 16 bytes			size 32
	c int32  //  8 bytes			size 40
	d int32  //  8 bytes			size 48 - Aligned on 8 bytes
}

func main() {
	// No byte padding
	var nbp nbp
	size := unsafe.Sizeof(nbp)
	// 无偏移量,连续的三个内存地址: sizeof[3]: 0xc0000200f0 0xc0000200f1 0xc0000200f2
	fmt.Printf("SizeOf[%d][%p %p %p]\n", size, &nbp.a, &nbp.b, &nbp.c)

	//Single byte padding,
	var sbp sbp
	size = unsafe.Sizeof(sbp)
	// 1byte 偏移量: SizeOf[4][0xc000020110 0xc000020112] 0[ 1 ]2
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &sbp.a, &sbp.b)

	// Three byte padding
	var tbp tbp
	size = unsafe.Sizeof(tbp)
	// 3byte 偏移量: SizeOf[8][0xc000020118 0xc00002011c] 8 [ 9 a b ]c
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &tbp.a, &tbp.b)

	// Seven byte padding
	var svnbp svnbp
	size = unsafe.Sizeof(svnbp)
	// 7byte 偏移量: SizeOf[16][0xc000020120 0xc000020128] 0 [1 2 3 4 5 6 7] 8
	fmt.Printf("SizeOf[%d][%p %p]\n", size, &svnbp.a, &svnbp.b)

	// Eight byte padding on 64bit Arch. Word size is 8 bytes.
	var ebp64 ebp64
	size = unsafe.Sizeof(ebp64)
	// 8 byte: SizeOf[48][0xc0000a0150 0xc0000a0160 0xc0000a0168 0xc0000a0178] 0xc0000a0160 0xc0000a0168 [0-8]
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &ebp64.a, &ebp64.b, &ebp64.c, &ebp64.d)

	// No padding.
	var np np
	size = unsafe.Sizeof(np)
	// SizeOf[40][0xc0000261b0 0xc0000261c0 0xc0000261d0 0xc0000261d4]
	fmt.Printf("SizeOf[%d][%p %p %p %p]\n", size, &np.a, &np.b, &np.c, &np.d)
}
