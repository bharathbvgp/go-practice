package main

import "fmt"

type ServerState int;

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// create a map of serverstate to string 

var stateName = map[ServerState] string {
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retry",
}

func (ss ServerState) String() string {
	return stateName[ss];
}

func transition(ss ServerState) ServerState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected , StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s" , ss));
	
	}
}



func main() {

	fmt.Println();
	fmt.Println(transition(StateIdle))
}