package main

import (
	"os"
	"imooc-demo/pipelineDemo/pipeline"
	"bufio"
	"fmt"
)

func main() {
	p := createPipeline(
		"small.in", 512, 4)
	writeToFile(p, "small.out")
	printFile("small.out")
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)

	for v := range p {
		fmt.Println(v)
	}

}
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	pipeline.WriterSink(writer, p)

}
func createPipeline(filename string, fileSize, ChunkCount int) <-chan int {
	chunkSize := fileSize / ChunkCount
	sortResults := []<-chan int{}
	for i := 0; i < ChunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}
