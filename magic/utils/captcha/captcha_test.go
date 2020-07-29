package captcha

import (
	"fmt"
	"os"
	"testing"
)

type byteCounter struct {
	n int64
}

func (bc *byteCounter) Write(b []byte) (int, error) {
	bc.n += int64(len(b))
	return len(b), nil
}

func BenchmarkNewImage(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(6)
	id := randomId()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewImage(id, d, StdWidth, StdHeight)
	}
}

func BenchmarkImageWriteTo(b *testing.B) {
	b.StopTimer()
	d := RandomDigits(6)
	id := randomId()
	b.StartTimer()
	counter := &byteCounter{}
	for i := 0; i < b.N; i++ {
		img := NewImage(id, d, StdWidth, StdHeight)
		img.WriteTo(counter)
		b.SetBytes(counter.n)
		counter.n = 0
	}
}

func BenchmarkImageWriteToXXX(b *testing.B) {
	d := RandomDigits(6)
	fmt.Println("*&*", string(d))
	id := randomId()
	img := NewImage(id, d, StdWidth, StdHeight)
	fmt.Println("^_^", id)
	// 写入到文件
	osfile, _ := os.Create("./test.png")
	img.WriteTo(osfile)
}
