# Docs for CLI

```plaintext
%> slipway --help

Usage: slipway sub-cmd

SubCommands:
    configuration
        add
        list
        remove
        export
    generate-configuration
        server
        client
    gossip-cluster
        list --watch
        remove <node>
        get-logs-from
    kubernetes-cluster
        initialize
        get-bootstrap-token
        kubeadm-action <node>
        replace-certificates
    version
```

```plaintext
slipway configuration add

Creating client configuration at ~/.slipway/config (from ${SLIPWAY_HOME} or ${HOME}/.slipway/)
Cluster Name: my-new-cluster
Cluster Endpoint: cluster-1.my-endpoint.domain
Cluster Port [7070]: 7070
Cluster API Secret: <16-bit key>
Cluster Custom CA File Path: (Default: ${SLIPWAY_HOME}/my-new-cluster/ca.pem if exists)
###############

slipway configuration list

my-new-cluster          (https://cluster-1.my-endpoint.domain:7070)
some-existing-cluster   (https://some.cluster.domain:7071)
```

```plaintext
slipway kubernetes initialize --help

Arguments
    cluster-name: Name of the slipway configuration (from ${SLIPWAY_HOME}/config)

Parameters
    --kubeadm-config path/to/kubeadm/config.yaml
    --etcd-ca-file
    --etcd-server-crt... (more of these???)


slipway kubernetes initialize my-new-cluster      # Defaults to ${SLIPWAY_CLUSTER} else you need to provide --cluster


Assessing ability to initialize Kubernetes Cluster
```