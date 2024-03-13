// The Manager is the other major component of our orchestration system.
// It will handle the following work:
// - Accept requests from users to start and stop tasks.
// - Schedule tasks onto worker machines.
// - Keep track of tasks, their states, and the machine on which they run.
package manager

import (
	"cubo/task"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// The Manager has queue, represented by the pending field,
// in which tasks will be placed upon first being submitted. The queue
// will allow the manager to handle tasks on a first-in-first-out (FIFO) basis.
// Next, the Manager has two in-memory databases:
// - one to store tasks and
// - another to store task events.
// The databases are maps of strings to Task and TaskEvent respectively
type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I will select an appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I will update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}
