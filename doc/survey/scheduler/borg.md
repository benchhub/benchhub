# Borg

Large-scale cluster management at Google with Borg

## Takeaway

- declarative configuration language BCL, a variant of GCL 
  - RCL https://github.com/at15/reika/issues/49
- store job, tasks, events in Dremel for analysis and capacity planning
- to reduce startup time, chose the machine that already have package installed
  
## 2. The user perspective

- workloads
  - long running *prod*
  - short running batch *non-prod*
- job
  - name, owner
  - tasks
  - soft & hard constraints
- task
  - linux container
  - resource requirement, cores, ram, disk, tcp ports
- BCL configuration language based on GCL
  - lambda functions
  - accumulated tens of millions of lines of BCL
  - similar to aurora
- alloc
  - a reserved set of resources on a machine in which one or more tasks can be run
- priority
  - monitoring, production, batch, best effort (testing or free)
- naming
  - BNS borg name service, a name for each task that is written into Chubby
- monitoring
  - http server that publish metrics
- record jobs, task events, as well as detailed pre-task resource usage information in Infrastore, based on Dremel

## 3. Borg architecture

- Borgmaster
  - store job and tasks in Paxos store
  - scheduler scan and pick the machine
    - feasibility checking
    - scoring, hybrid EPVM + best fit
    - to reduce startup time, the scheduler prefers to assign tasks to machines that already have the necessary packages 
- Borglet
  - Borgmaster polls Borglet every few seconds 