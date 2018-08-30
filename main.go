package main

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin/binding/example"
	"github.com/golang/protobuf/proto"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = ioutil.WriteFile("./dat1", data, 0644)
	check(err)

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// etc.
}
