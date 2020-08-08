# v0.2.0 Runtime

## TODO

- [ ] split into small features that can work in parallel w/ v0.1.0
- [ ] resource counting, runtime reports back to core?

## Overview

Provide a managed runtime so user only need to provide code.

## Related

- Previous: [v0.1.0 Micro](../v0.1.0-micro)

## Motivation

Allow running test and benchmarks locally and in the cloud. Running locally is useful for interactive development.
Running in the cloud is for CI and large scale benchmarks. The cloud version should also focus on multi-tenancy and efficient scheduling.

## Specs

There are two scenarios for running benchmarks, local or remote.

Local

- single user
- triggered manually
- small test/benchmark, won't eat all the resources of a mbp
- take a snapshot of local code, so it's possible to compare different branches or even same commit but different parameters

Remote

- multiple users from multiple organizations
- triggered by webhook like git commit, bot command
- small, big, requires multiple nodes and may contain workflow (pipeline, trigger a large test if small test pasts)
- binpack and pick the cheapest offer to reduce cost on cloud service providers

Proposed examples are:

Local Examples

- sort
- strstr
- loc

Remote Examples

- xephon-b
- tpc
  - can use pingcap/go-tpc
  
## Background

In the very first benchhub implementation, I tried to write both benchhub and container orchestration (e.g. k8s) at same time.
The easier approach would be use existing platforms, e.g. CRD in k8s.

Using k8s has the following benefits:

- faster prototype
  - the eco system is large and provide metric, log etc. out of box
- deploy to different cloud provider or on premise cloud is easier because they all talk in YAML

Using k8s has the following drawbacks:

- waste resource, k8s's own scheduler and autoscaler maybe not be that good for benchmark type of batch jobs
- overhead and needless complication due to using container (and container network, storage etc.)

Using VM has the following benefits:

- more knobs to tune?
  - might able to do the same using container ...?
- can use bare metal on providers that provide them
  - this is good for database because you don't binpack databases in performance test
  - it should also be possible for k8s by using label and customize the cloud controller logic?
- for a single vm with a single database, it might be faster to download binary compared w/ pull container, setup container network etc.

Using VM has the following drawbacks:

- deal with different cloud provider API and SDKs
  - should either use interface or grpc to write plugin (latter is similar to terraform) 
  
## Features

### v0.1.1 local simple batch

Description

Run a batch job locally using docker? (It might be the easies way to package code ...)
Collect metrics, result, log and report to local benchhub instance.
Each framework can run as a host process and use docker log API and bind mount to access output and data.

It is similar to exp/qaq16, though we may want to use docker API instead of shell out.
Well ... honestly I am not a big fan of docker API client and there were several attempts to write a new one in go.ice.

Components

- `core/runtime`
  - allow user to submit request and use a managed runtime to run a job
  - define the interface for runtime to implement
  - runtime should be able to
    - build the code (container)
    - run the code (container)
- `runtimes/local`
  - the control plane as a process/go routine
  - package code to a container
  - run the container
  - call framework to analysis output
- `lib/fastmonitor`
  - fetch container metrics using cgroup w/o using the slow docker API

### v0.1.2 k8s simple batch

Description

Run a batch job using a single pod on k8s. Collect metrics, result, log and report to benchhub.
Each framework should have its own docker image that pulls the code, run test/benchmark, collect the output.
The draw back it runs multiple processes in one container and if the test code OOM, the result is gone.
Streaming the result or run the framework as a sidecar container might help.

We also need to consider where to build the container. For language specific benchmark, we can prebuild the images,
and build the code inside those containers. For database benchmark, might need a builder to build database binary,
and another runner image that contains the binary. Though it's possible to build and run database in same image.

Components

- `runtimes/k8s`
  - the control plane as a k8s CRD (it's possible to run it in process w/ proper cert)
- `runtimes/k8s/eks`
- `runtimes/k8s/gke`
- `lib/gogocloud/aws`
  - handy wrapper for aws sdk
- `lib/gogocloud/gcp`
  - handy wrapper for gcp sdk

### v0.1.3 vm simple batch

Description

Allocate a new vm for each job. Both framework and test run in the VM as container or process?
This is a waste of resource for sure, but I want to delay scheduler until more complex workflow is introduced.

Components

- `runtimes/remote`
- `runtimes/remote/aws`
- `runtimes/remote/ali`
- `runtimes/remote/gcp`
