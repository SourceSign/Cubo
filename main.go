package main

import (
	"cubo/manager"
	"cubo/node"
	"cubo/task"
	"cubo/worker"
	"fmt"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	t := task.Task{
		ID:     uuid.New(),
		Name:   "task-1",
		State:  task.Pending,
		Image:  "Image.1",
		Memory: 1024,
		Disk:   1,
	}
	te := task.TaskEvent{
		ID:          uuid.New(),
		TargetState: task.Pending,
		Timestamp:   time.Now(),
		Task:        t,
	}

	w := worker.Worker{
		Name:  "worker-1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	m := manager.Manager{
		Pending: *queue.New(),
		TaskDb:  make(map[string][]*task.Task),
		EventDb: make(map[string][]*task.TaskEvent),
		Workers: []string{w.Name},
	}

	n := node.Node{
		Name:          "Node-1",
		IPAddress:     "192.168.1.1",
		Cores:         4,
		MaximumMemory: 1024,
		MaximumDisk:   25,
		Role:          "worker",
	}

	fmt.Printf("task: %v\n", t)
	fmt.Printf("taskEvent: %v\n", te)
	fmt.Printf("worker: %v\n", w)

	w.CollectStats()
	w.RunTask()
	w.StartTask()
	w.StopTask()

	fmt.Printf("manager: %v\n", m)

	m.SelectWorker()
	m.UpdateTasks()
	m.SendWork()

	fmt.Printf("node: %v\n", n)

}
