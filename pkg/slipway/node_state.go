package slipway

// NodeState items are the states a cluster can be in and is directly related to the unready/bootstrap states of kubernetes
type NodeState int

func (n NodeState) String() string {
	switch n {
	case NodeWaitingToBootstrap:
		return "waiting-to-bootstrap"
	case NodeBootstrapping:
		return "bootstrapping"
	case NodeReady:
		return "ready"
	case NodeFailedBootstrap:
		return "busy"
	default:
		panic("Didn't find expected constant for NodeStates")
	}
}

const (
	// NodeWaitingToBootstrap state signifies node has started, and no one in the cluster is in the Ready State
	// This is the startip initial state
	NodeWaitingToBootstrap NodeState = iota

	// NodeBootstrapping state means the node is performing actions to bootstrap itself
	NodeBootstrapping

	// NodeReady state means the node is ready to help bootstrap other nodes
	NodeReady

	// NodeFailedBootstrap state is a holding state for Failed Bootstrap events
	NodeFailedBootstrap
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

	State: FailedBootstrap is for when kubeadm fails for some reason and this state needs manual intervention
	- wait for user action to either "force-re-bootstrap" or die or something
*/
