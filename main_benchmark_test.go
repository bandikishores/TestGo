package main

import (
	"testing"

	pb "bandi.com/main/pkg/data"
)

func BenchmarkTest(t *testing.B) {
	for n := 0; n < t.N; n++ {
		PerformProtoChanges(&pb.AddressBook{})
	}
}
