package main

import (
	"encoding/json"
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
			{ "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" },
			{ "url": "dummy", "type": "anotherType" }
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

type Avatar struct {
	URL     string
	NewType string
}

func TestJeffailGabsParser(t *testing.T) {
	jsonParsed, err := gabs.ParseJSON([]byte(`{"name":{"first":"kishore"}}`))
	if err != nil {
		panic(err)
	}

	container := jsonParsed.S("name", "first")
	container.Set("Modified")
	fmt.Printf("Modified Using First Container JSON : %s\n", jsonParsed.String())

	container = jsonParsed.S("name")
	container.Set("Modified", "first")
	fmt.Printf("Modified Using Name Container JSON : %s\n", jsonParsed.String())
	if true {
		return
	}
	jsonParsed, err := gabs.ParseJSON(jsonByteGlobal)
	if err != nil {
		panic(err)
	}

	container := jsonParsed.S("person", "name", "first")
	_, err = container.Set("modifiedFirstBuhahahaha")
	jsonParsed.S("person", "name").Set("domeDummy")
	// jsonParsed.Set("person", "name", "first", container)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Container Name : %s\n", jsonParsed.String())

	for name, children := range jsonParsed.Search("person").ChildrenMap() {
		fmt.Printf("Array Child : %s, value : %s\n", name, children.String())
		switch children.Data().(type) {
		case string:
			fmt.Printf(" just got string\n")
		case []interface{}:
			fmt.Printf(" got array interface\n")
			for _, children1 := range children.Children() {
				fmt.Printf(" got recursive array of %s\n", children1.Data())
			}
		case map[string]interface{}:
			fmt.Printf(" got map this is nested oject\n")
		default:
			fmt.Printf(" nt handling \n")
		}

	}

	for name, children := range jsonParsed.ChildrenMap() {
		switch jsonParsed.Data().(type) {
		case json.Number:
			val, err := jsonParsed.Data().(json.Number).Int64()
			if err != nil {
				fmt.Errorf("Error %v", err)
			} else {
				fmt.Printf("Int Type : %d", val)
			}

		case json.RawMessage:
			val, err := jsonParsed.Data().(json.Number).Int64()
			if err != nil {
				fmt.Errorf("Error %v", err)
			} else {
				fmt.Printf("Raw Type : %d", val)
			}
		}
		fmt.Printf("First Child : %s, value : %s\n", name, children.String())
	}

	fmt.Printf("Full Name : %s\n", jsonParsed.S("person", "name", "fullName").String())

	fmt.Printf("Before Modification : %s\n", jsonParsed.String())
	//jsonParsed.Delete("person", "name", "fullName")
	jsonParsed.Set("Kishore Bandi", "person", "name", "fullName")
	jsonParsed.Set(map[string]struct{}{}, "previous")
	fmt.Printf("After Modification : %s\n", jsonParsed.String())

	// S is shorthand for Search
	for key, child := range jsonParsed.S("person").ChildrenMap() {
		// fmt.Printf("key: %v, value: %v\n", key, child.Data().(string))
		if key == "avatars" {
			for _, avatar := range child.Children() {
				fmt.Println(avatar.S("url").String())
				avatar.Set("Added To Existing", "NewKey")
			}
			err := child.ArrayAppendP(Avatar{URL: "NewURL", NewType: "NewFieldAsWell"}, "0")
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("After Modification : %s\n", jsonParsed.String())
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
