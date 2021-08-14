package main

import (
	enumpb "github.com/SarathLUN/udemy-protocol-buffers-3/my-code/06-golang/02/src/enum"
	"github.com/SarathLUN/udemy-protocol-buffers-3/my-code/06-golang/02/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {

	//sm := doSimple()
	// working with file
	//readAndWriteDemo(sm)

	// working with json
	//jsonDemo(sm)

	// working with enum
	doEnum()

}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           44,
		DayOfTheWeek: enumpb.DayOfTheWeek_MONDAY,
	}
	log.Println(em)

	// re-assign value
	em.DayOfTheWeek = enumpb.DayOfTheWeek_SUNDAY
	log.Println(em)

}

func jsonDemo(sm proto.Message) {
	// convert to JSON
	smAsString := toJSON(sm)
	log.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	log.Println(sm2)
}

func toJSON(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON:", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Panic(err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	// write to file
	err := writeToFile("./my-code/06-golang/02/out/simple.bin", sm)
	if err != nil {
		log.Panic(err)
	}

	// read from file
	// create empty struct
	sm2 := &simplepb.SimpleMessage{}
	err = readFromFile("./my-code/06-golang/02/out/simple.bin", sm2)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Read the content:", sm2)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(in, pb)
	if err != nil {
		panic(err)
	}
	return nil
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
