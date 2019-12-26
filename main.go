package main

import (
	"bandi.com/main/data"
	"bandi.com/main/state"
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

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
