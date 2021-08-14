package main

import (
	"github.com/SarathLUN/udemy-protocol-buffers-3/my-code/06-golang/02/src/simple"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()
	err := writeToFile("./my-code/06-golang/02/out/simple.bin", sm)
	if err != nil {
		log.Panic(err)
	}
	//readFromFile()
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes:", err)
	}
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file:", err)
		return err
	}
	log.Println("Data has been written!")
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3},
	}
	log.Println(sm)

	sm.Name = "I rename you"
	log.Println(sm)

	log.Println("The ID is:", sm.GetId())

	return &sm
}
