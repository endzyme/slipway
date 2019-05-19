# Next
1. Wire in the local kubernetes status detection to the lifecycle-state of the slipway cluster
    - initial thoughts: how do we know if kubernetes is running, what files should be there and should not be there? how would we manually change state if we needed to?
2. Wire in lifecycle state watcher loops
    - if i'm waiting-to-bootstrap then always be looking for ready nodes in the gossip channel with role "master"
    - if i'm in the bootstrapping state then i should just be in a holding pattern until someone moves me to the ready state or broken state
    - if i'm ready ... mmm not sure what to watch for here? maybe command that tell you you're not ready anymore?
3. Wire in the first initialization command with kubeadm to the gRPC server
    - step 1 - call a single sever out and say "HEY YOU BOOTSTRAP YO SELF" and stream all kubeadm logs back to cli client with status of commands

## Mental notes
Components so far:
- cluster: responsible for gossip cluster connection and starting all the cluster lifecyclestate watchers, also responsible for managing tags of the cluster based on local node state changes (logic in node component)
- kubeadm: responsible for handling kubeadm calls gracefully and being able to stream buffers out
- node: All code responsible for detecting node state that affects lifecyclestate of clustermember
- cmd server/client: responsible for starting the server with the right context and waiting for syscall signals