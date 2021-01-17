package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// Task definitions
type MapTask struct {
	Filename string
	ID       int
	NReduce  int
}

type ReduceTask struct {
	ID   int
	NMap int
}

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

type GetTaskArgs struct{}

type GetTaskReply struct {
	// Assigned Task
	IsReduce bool
	RTask    *ReduceTask
	MTask    *MapTask
}

type PostCompletionArgs struct{}

type PostCompletionReply struct{}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
