package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

// 数组数据源节点
func ArraySource(a ...int) <-chan int {
	out := make (chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 内部排序节点
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// Read into memroy
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 归并节点
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
	}()

	return out
}

// 读取节点
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)

	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				break
			}
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
		}
		close(out)
	}()
	
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}