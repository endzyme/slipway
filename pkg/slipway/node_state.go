package slipway

// NodeState items are the states a cluster can be in and is directly related to the unready/bootstrap states of kubernetes
type NodeState int

const (
	// WaitingToBootstrap state signifies node has started, and no one in the cluster is in the Ready State
	// This is the startip initial state
	WaitingToBootstrap NodeState = iota

	// Bootstrapping state means the node is performing actions to bootstrap itself
	Bootstrapping

	// Ready state means the node is ready to help bootstrap other nodes
	Ready

	// Busy state is for nodes that are performing an action that should not be interrupted
	Busy
)

/*
  State: WaitingToBootstrap
    - Start and look for Ready nodes
      - None Found: WaitForReadyNodes(timeout time.Duration)
      - Found: AskForBootstrapDonation(nodeRole [master|node]) => Transition to Bootstrapping State

  State: Bootstrapping timeout after 300s
	- Call a node that is able to donate
	  - No: Try again
	  - Yes: synchronously wait for bootstrap token
	- Join the cluster with bootstrap token and synchronously wait for kubeadm call to pass
	  - Pass: transition to ready
	  - Fail: ... not sure yet

  State: Ready (just wait for users or the system)
	- continually watch the health of kube or sometihng
	- wait for a user to transition the state manually

  State:
*/
