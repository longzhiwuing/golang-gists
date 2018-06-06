package main

import (
	"imooc-demo/pipelineDemo/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	//mergeDemo()

	const filename  = "large.in"
	const n = 100000000

	file, err := os.Create(filename)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()
	file, err = os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file))

	count:=0
	for v := range p {
		fmt.Println(v)
		count++
		if count>= 100 {
			break
		}

	}

}
func mergeDemo() {
	//p := pipeline.ArraySource(3, 2, 6, 7, 4)
	/*for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}*/
	//p := pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4))
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),

		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	for v := range p {
		fmt.Println(v)
	}
}
