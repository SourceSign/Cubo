// The task package implements the task which is the smallest unit of work
// in the cube system and typically runs in a container.
package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// The State type represents the states a task goes through,
// from Pending, Scheduled, Running, to Failed or Completed.
//
//	+---------+       +-----------+         +---------+
//	| Pending |  ---> | Scheduled |  -----> | Running |
//	+---------+       +-----------+         +---------+
//	   |                                         |
//	   |                                         V
//	   |               +--------+          +-----------+
//	   +-------------> | Failed |  ------> | Completed |
//	                   +--------+          +-----------+
type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

// Task struct which ist the smallest unit of work with Podman-specific fields as
// cubo is limited to dealing with Podman containers.
type Task struct {
	ID            uuid.UUID // uniquely identify individual tasks
	ContainerID   string    // This ID will allow us to interact with the running task.
	Name          string    // human-readable Name
	State         State     // Indicate the current state the task
	Image         string    // Docker image a task should use
	CPU           float64
	Memory        int64 // amount of memory a task needs
	Disk          int64 // amount of diks space a task needs
	ExposedPorts  nat.PortSet
	PortBinding   map[string]string // used by Docker to ensure the machine allocates the proper network ports for the task
	RestartPolicy string            //  tell the system what to do in the event a task stops or fails unexpectedly.
	StartTime     time.Time         // Timestamp when a task starts
	FinishTime    time.Time         // Timestamp when a task stops
}

// The TaskEvent struct, which represent an event that moves a Task from one state to another.
// internal object that our system uses to trigger tasks from one state to another.
type TaskEvent struct {
	ID          uuid.UUID
	TargetState State     // Indicate the state the task should transition to (e.g. from Running to Completed).
	Timestamp   time.Time // timestamp to record the time the event was requested.
	Task        Task      // task that is triggerd from one state to another
}
