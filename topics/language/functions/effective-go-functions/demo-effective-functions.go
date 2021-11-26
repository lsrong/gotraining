package effective_go_functions

import (
	"io"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

// NextInt 获取字节切片中的数字,以及下一个的下标
func NextInt(b []byte, pos int) (value, nextPos int) {
	for ; pos < len(b) && !isDigit(b[pos]); pos++ {
	}
	for ; pos < len(b) && isDigit(b[pos]); pos++ {
		value = value*10 + int(b[pos]) - '0'
	}

	return
}

// ReadFull  Named result parameters
func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
