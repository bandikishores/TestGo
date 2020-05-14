package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"bandi.com/main/data"
	pb "bandi.com/main/pkg/data"
	"bandi.com/main/pkg/protocol"
	"bandi.com/main/state"
	"github.com/golang/protobuf/proto"
)

var channel chan struct{}

// ResolvedAddresses of host.
func getResolvedAddresses(host string) *net.TCPAddr {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		log.Fatalln("ResolveTCPAddr of host:", err)
	}
	return addr
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write memory profile to file")
	flag.Parse()

	if cpuprofile != nil && *cpuprofile != "" {
		fmt.Println("cpu profiler enabled")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	} else {
		fmt.Println("cpu profiler disabled")
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
	fmt.Println("mem profiler disabled")

	go protocol.RunRestServer(ctx, "18091", "18081")
	protocol.RunGrpcServer(ctx, "18091")

	// PerformProtoChanges(nil)

	// TestContext(ctx)

}

// TestGoRoutine - Sample Test
func TestGoRoutine() {
	channel = make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			channel <- struct{}{}
			log.Printf("Sent")
			channel <- struct{}{}
			log.Printf("Sent")
			channel <- struct{}{}
			log.Printf("Sent")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-channel
			log.Printf("Got")
		}
	}()

	wg.Wait()

	if true {
		return
	}
}

// TestContext - Some test
func TestContext(ctx context.Context) {
	session := state.SessionInfo{
		SessionID: "123",
	}

	var restoredSession *state.SessionInfo

	newContext := context.WithValue(ctx, data.Session, &session)
	restoredSession = newContext.Value(data.Session).(*state.SessionInfo)

	fmt.Println("Value of Session ", session)
	fmt.Println("Value of Restored Session ", restoredSession)

	restoredSession.SessionID = "456"

	fmt.Println("Value of Session ", session)
	fmt.Println("Value of Restored Session ", restoredSession)

	change(session)

	fmt.Println("Value of Session ", session)
	fmt.Println("Value of Restored Session ", restoredSession)

	var sdf string
	value := sdf
	fmt.Println("Value of sdf ", sdf, value)
}

// PerformProtoChanges - Proto Buf Serialization & DeSerialization
func PerformProtoChanges(book *pb.AddressBook) *pb.AddressBook {
	if book == nil {
		book = &pb.AddressBook{People: []*pb.Person{
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
	}
	// ...

	fname := "/tmp/prototest.out"

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	restoredBook := &pb.AddressBook{}
	if err := proto.Unmarshal(in, restoredBook); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	return restoredBook
}

func change(session state.SessionInfo) {
	session.SessionID = "789"
}

func checkType(typee interface{}) string {
	switch typee.(type) {
	case int, int64:
		return "integer"
	default:
		return "unknown"
	}
}

// DummyFunc - Dummy Function
func DummyFunc() string {
	return "dummy"
}
