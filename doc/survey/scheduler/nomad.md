# Nomad

https://github.com/hashicorp/nomad

- https://www.nomadproject.io/docs/internals/scheduling.html
  - create job
  - enter a queue
  - evaluation
  - allocation plan
  - ranking
  
- job
  - group, all the task in one group is assigned to same node
    - task
    
- nomad/nomad/worker.go 
  - `func (w *Worker) run()` the long-lived goroutine which is used to run the worker
- nomad/nomad/heartbeat.go