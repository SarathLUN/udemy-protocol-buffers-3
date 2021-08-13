package main

import (
	"github.com/SarathLUN/udemy-protocol-buffers-3/my-code/06-golang/01/example/simple"
	"log"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 23, 45, 6},
	}
	log.Println(sm)
}
