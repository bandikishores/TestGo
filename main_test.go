package main

import (
	"testing"

	pb "bandi.com/main/pkg/data"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

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
