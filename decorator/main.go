package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func CapReader(r io.Reader) io.Reader { return &capitalizedReader{r: r} }

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	// 此处调用了原生的 Read 方法
	n, err := r.r.Read(p)

	if err != nil {
		return 0, err
	}
	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}
func main() {
	r := strings.NewReader("hello, gopher!\n")
	r1 := CapReader(io.LimitReader(r, 4))
	// stdout：就是正常的终端打印输出，比如你在代码中使用log或者fmt输出了一句话，就是通过这个流来处理。
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
