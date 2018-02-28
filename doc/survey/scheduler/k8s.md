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

## Nodes

- https://kubernetes.io/docs/concepts/architecture/nodes/
  - register in master by human, master will try to reach the node, i.e. using its IP or FQDN provided in meta as name
  - let node register itself to master when start
- node status
  - addresses, HostName, External IP, Internal IP
  - condition, OutOfDisk, Ready, MemoryPressure, DiskPressure, NetworkUnavailable, ConfigOK
  - capacity, CPU, memory, maximum number of pods
  - info, kernel version, k8s version, docker version, os name

````json
{
  "kind": "Node",
  "apiVersion": "v1",
  "metadata": {
    "name": "10.240.79.157",
    "labels": {
      "name": "my-first-k8s-node"
    }
  }
}
````

- https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
- https://github.com/kubernetes/node-problem-detector
  - used to detect issues like kernel deadlock, corrupted fs, broken container runtime
  - seems to have low overhead https://github.com/kubernetes/node-problem-detector/issues/2#issuecomment-220255629

## Master Node Communication

- https://kubernetes.io/docs/concepts/architecture/master-node-communication/
- apiserver -> kublet
  - fetching logs for pods
  - attaching through kubectl to running pods
  - providing the kubelet's port-forwarding functionality
- apiserver -> nodes, pods and services

## Pod

- https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/

````yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo Hello Kubernetes! && sleep 3600']
````

## Job

- https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/

````yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4
````
