# Kubernetes

- https://github.com/benchhub/benchhub/issues/3
  - https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html
- https://blog.heptio.com/core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca
  - the flow image is very clear
  - use watch in etcd to coordinate between components, like pubsub, but the topic is rich instead of message
    - client mirror a subset in memory, watch keep cache up to date, fall back to polling if watch failed
   - api server talks to etcd directly, other components talk to api server
- https://dzone.com/articles/kubernetes-lifecycle-of-a-pod
  - readiness and liveness check 
    - https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
- [kubernetes设计理念分析 | 从运行流程和list-watch看kubernetes系统的设计理念](https://mp.weixin.qq.com/s?__biz=MzUzNzYxNjAzMg==&mid=2247483683&idx=1&sn=0a92e0e8e0d54d6ee76852f9db45b181&pass_ticket=HZPrXHGPCWrYW4HMZTV9MZ19kAGoK72zI%2FlkaARMhlOZLmeOGGuGDjrhWBOGbSj2)
  - use `level trigger` instead of `edge trigger` to avoid using external mq like openstack (rabbitmq), cloudFoundary (nats)
    - https://stackoverflow.com/questions/1966863/level-vs-edge-trigger-network-event-mechanisms
  - use http/2 streaming for watching
  - `ResourceVersion` use etcd to have global increasing key, client provide cache version and server push all newer versions to client
  - before `watch` use `list` to get global status and version, then watch
    - if error when `watch`, use `list`

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
