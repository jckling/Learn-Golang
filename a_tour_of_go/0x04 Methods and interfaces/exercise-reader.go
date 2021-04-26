package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// 产生一个 ASCII 字符 'A' 的无限流
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
