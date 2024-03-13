// The worker package implements the next layer that sits atop the foundation. It's main purpose is to keep track of tasks.
package worker

import (
	"cubo/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// The worker struct which is doing the following:
//
//   - Run tasks as Docker containers.
//   - Accept tasks to run from a manager.
//   - Provide relevant statistics to the manager for the purpose of scheduling tasks.
//   - Keep track of its tasks and their state.
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

// Function which can be used to periodically collect statistics about the worker.
func (w *Worker) CollectStats() {
	fmt.Println("I will collect stats!")
}

// Function that handles running a task on the machine where the worker is running.
func (w *Worker) RunTask() {
	fmt.Println("I will run or stop tasks!")
}

// Function that starts a task
func (w *Worker) StartTask() {
	fmt.Println("I will start a task!")
}

// Function that stops a task
func (w *Worker) StopTask() {
	fmt.Println("I will stop a task!")
}
