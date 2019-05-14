# TODO
* Build a CLI client that connects to the http endpoint
* Configure some hard set "actions" for bootstrapping a node
* Configure detection of kubernetes cluster state on startup (already running, not running, etc)
* Add default labels to nodes like "master" vs "node" type metadata
* Add labels to nodes which include IP and hostname
* Configure the run-actions to be pushed over serf's event "bus"
* Make considerations for TLS on gossip channel
* Make considerations for TLS on http channel

## Confiuration
* Have CLI Client read from a ~/.telephone config file (maybe yaml??)
* Have the server side app read from environment variables and a config file if needed
* Minimum requirements for config are where to cluster (address elb etc)
* Configure a preshared key in the CLI and Serverside configs for auth between API and gossip commands

## Packaging and Installing
* Server side should be installable with deb or a docker container or maybe rpm??? or just run the binary

## Consider User Experience
* I want to be able to install the serverside binary on 3 masters in an autoscaling group with an ELB for gossip discovery
* They should all come up and see each othe as members
* The CLI tool should be able to connect to another ELB to fire rpc commands to the nodes
* We should support actions that have "broadcast" vs "directed" messages
  * Example broadcast message may be: "Install all the needed elements for Kubernetes v1.13.x on all Masters"
  * Example directed command: "Tell node-abcd to bootstrap as a master with kubeadm and when complete fire join action to awaiting masters"
* MVP maybe we should support providing your own CA and other bootstrapping material as inputs to CLI tool
  * Example would be PKI infra or the KubeADM API docs
* We should be able to BLOCK "download k8s components for install" if cluster is running state
  * Nodes may need a state setting to allow or disallow certain cluster actions
  * Example: Don't allow bootstrap with kubeadm if cluster is already bootstrapped
* Failed executions should be as idempotent as possible when transitioning states

# Major Components
1. Gossip Clustering
2. API Server & Client CLI
3. Kubeadm Command Wrappers (versioned)

# Workflow for initialization
1. Three servers come up and join each other

# Requirements before Bootstrapping
1. Has Kubernetes installed (kubelet, kubeadm, etc)
2. Has Slipway Server Installed and Configured
3.


# Server Configuration
- Gossip Discovery Listening address and port (default 0.0.0.0:6776/udp)
- API Server Listening Address and Port (default 0.0.0.0:7070/tcp)
- Certificates

# CLI Features
- Generate a config file for the Server
- Download the server binary matching your version of the CLI
- Generate (or provide docs on) a self-signed long-lived certificate for gossip mtls