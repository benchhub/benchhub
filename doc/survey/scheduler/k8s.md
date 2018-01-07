# Kubernetes

https://github.com/benchhub/benchhub/issues/3

## Architecture

https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/architecture.md

- Cluster control plane (master)
  - API server
  - Cluster state store `etcd`
  - Controller Manager, make sure the desired state is met?
  - Scheduler, two level?
- The Kubernetes Node
  - Kublet
    - execute pods using container
    - use cAdvisor for resource monitoring
  - Container runtime
    - docker
    - rkt
    - [cri-o](https://github.com/kubernetes-incubator/cri-o)
  - Kube Proxy
    - ? implementation?
    - 'Each node runs a kube-proxy process which programs iptables rules to trap access to service IPs and redirect them to the correct backends'