// The Node component represents a physical machine where the worker and tasks will run.
package node

// A node is an object that represents any machine in our cluster.
type Node struct {
	Name            string //
	IPAddress       string // IP address, which the manager will want to know in order to send tasks to it
	Cores           int    // number of cores on the node
	MaximumMemory   int    // maximum amount of Memory of the physical machine
	MaximumDisk     int    // maximum amount of Disk space of the physical machine
	AllocatedMemory int    // allocated amount of memory used by the tasks on a machine
	AllocatedDisk   int    // amount of  disk used by the tsks on a machine
	Role            string // the roles of the, ether manager and worker“
	TaskCount       int    // Number of tasks on the node
}
