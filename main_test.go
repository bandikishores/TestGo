package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	pb "bandi.com/main/pkg/data"
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
