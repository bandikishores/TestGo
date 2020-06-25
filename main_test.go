package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Jeffail/gabs/v2"
	"github.com/buger/jsonparser"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	pb "bandi.com/TestGo/pkg/data"
)

var once sync.Once
var globalString string
var i int32 = 0

func TestMain(t *testing.T) {
	want := "dummy"
	assert.Equal(t, want, DummyFunc())
}

func TestProto(t *testing.T) {
	book := &pb.AddressBook{People: []*pb.Person{
		{
			Name:  "Kishore Bandi",
			Id:    123,
			Email: "bandikishores@gmail.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "124", Type: pb.Person_HOME},
				{Number: "125", Type: pb.Person_MOBILE},
			},
		},
	}}
	restoredBook := PerformProtoChanges(book)
	assert.True(t, proto.Equal(book, restoredBook))
}

var jsonByteGlobal []byte = []byte(
	`{
		"person": {
		  "name": {
			"first": "Leonid",
			"last": "Bugaev",
			"fullName": "Leonid Bugaev"
		  },
		  "github": {
			"handle": "buger",
			"followers": 109
		  },
		  "avatars": [
			{ "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
		  ]
		},
		"company": {
		  "name": "Acme"
		}
	  }`)

func TestBugerJsonParser(t *testing.T) {
	jsonByteArray := jsonByteGlobal

	// You can specify key path by providing arguments to Get function
	value, dataType, _, err := jsonparser.Get(jsonByteArray, "person", "name", "fullName")
	if err != nil {
		fmt.Printf("errors was : %v", err)
	}
	switch dataType {
	case jsonparser.String:
		fmt.Printf("string %s", value)
	case jsonparser.Number:
		fmt.Printf("Number %d", value)
	default:
		fmt.Printf("Unknown %v", value)
	}

	jsonparser.ArrayEach(jsonByteArray, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(jsonparser.Get(value, "url"))
	}, "person", "avatars")

	json, err := jsonparser.ParseString(jsonByteArray)
	if err != nil {
		fmt.Printf("errors was : %v", err)
	}
	fmt.Printf("Before Modification : %s", json)

	jsonByteArray = jsonparser.Delete(jsonByteArray, "person", "name", "fullName")
	json, err = jsonparser.ParseString(jsonByteArray)
	if err != nil {
		fmt.Printf("errors was : %v", err)
	}
	fmt.Printf("After Deletion : %s", json)

	jsonparser.Set(jsonByteArray, []byte("Kishore Bandi"), "person", "name", "fullName")
	json, err = jsonparser.ParseString(jsonByteArray)
	if err != nil {
		fmt.Printf("errors was : %v", err)
	}
	fmt.Printf("After Modification : %s", json)
	fmt.Printf("Original bytearray : %s", string(jsonByteArray))
}

func TestJeffailGabsParser(t *testing.T) {
	jsonParsed, err := gabs.ParseJSON(jsonByteGlobal)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Full Name : %s", jsonParsed.S("person", "name", "fullName").String())

	fmt.Printf("Before Modification : %s", jsonParsed.String())
	//jsonParsed.Delete("person", "name", "fullName")
	jsonParsed.Set("Kishore Bandi", "person", "name", "fullName")
	fmt.Printf("After Modification : %s", jsonParsed.String())
	fmt.Printf("Original bytearray : %s", string(jsonByteGlobal))
}

func TestGrpcProto(t *testing.T) {
}

func GetString() string {
	if globalString == "" {
		once.Do(func() {
			if i == 0 {
				i++
				return
			}
			id, err := uuid.NewUUID()
			if err != nil {
				panic(fmt.Sprintf("Error occurred while creating UUID %v", err))
			}
			globalString = id.String()
		})
	}
	return globalString
}
