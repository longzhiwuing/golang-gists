package main

type intGen func() int

func (intGen) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func main() {
}
