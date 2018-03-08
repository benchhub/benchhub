# Store

- issues
  - Rethink node, job https://github.com/benchhub/benchhub/issues/30
  - Job status stages https://github.com/benchhub/benchhub/issues/27

- meta store is mainly used to keep state of the cluster, which including two part
  - node states
  - job states
- additionally, all the update to state should be logged
  - node register history
  - job update history
  
currently we just store all the meta inside central's memory, agent also need store due to long running jobs

- [ ] might add a driver called kill, which stop specific task in long running stages

## Node store

- node ~~config~~ info
  - ip etc.
  - [ ] a message for address, bindAddr etc.
  - provider
  - capacity
  - preferred role
  - state
- node status
  - state, idle, job running, job waiting (peers), report, clean up
  - jobs
    - foreground stage
      - tasks
        - status
    - background stage
      - tasks
        - status
        
        
````text
[
    id string # same as info.id
    state enum # same as status.state
    # provided when node register
    info: {
        id string # generated when node start, central also use it instead of assign new one
        host string # hostname
        addr {
            ip # set manually or obtained from remoteAddr, the ip address other people can reach
            bindAddr
            remoteAddr # filled by other side, what they see when they receive the request
        }
        startTime int64 # agent/central process start time unix second
        bootTime int64 # machine boot time from /proc/stat
        role # preferred role
        provider # iaas, packet, aws etc.
        capacity # node capacity when it register
    },
    status: {
        state enum
        // TODO: job info
    }
    jobs: [
    // TODO: running and longRunning
    ]
]
````
    
## Job store

- job spec
- nodes assignment
- current stage
  - status of all nodes
- [ ] TODO: copy is needed between config struct and proto message
- ref https://developer.github.com/v3/repos/#get

spec

````text
{
    id string # assigned when job is created
    name string 
    Owner: {
        id string
        name string
        type enum # user, org
    }
    Workload: {
        framework string
        frameworkVersion string
        database string
        databaseVersion string        
    }
    NodeAssignments: [
        {
            name string # the unique name used throughout config TODO: might put it into dns?
            role enum # node role, loader or database
            # TODO: add label to nodes
            # TODO: select node etc. by resource etc.
            selectors [
            ]
        }
    ]
    Pipelines: [
        {
            name string # unique in pipelines
            stages []string # stages that can run in parallel
        }
    ]
    Stages: [
        {
            name string # unique name used throughout config
            type enum # short running (default), long running (must have at lease one long running task in this stage), stopper (special stage used to stop long running stage)
            selectors [
                # OR relation between multiple selectors
                {
                # TODO: use label when it is supported
                    name string,
                    role enum
                }
            ]
            # optional, used for running tasks in parallel, if used, MUST specify all the tasks, and all the tasks MUST have name
            pipelines [
                {
                    name string # unique in pipelines
                    tasks []string # tasks that can run in parallel
                }
            ]
            tasks [
                {
                   name string # unique in stage MUST be specified if pipeline is used or is longRunning
                   longRunning bool
                   driver enum # exec, shell, docker
                   env: [
                      {
                        k string
                        v string
                      }
                   ]
                   // TODO: does node selector apply to stopper? yes! 
                   stopper: {
                       stage string
                       task string
                       all bool # stop all long running tasks in a stage
                   }
                   shell: {
                       command string # passed to sh -c                        
                   }
                   exec: {
                       command string # path to binary
                       args []string
                   }
                   docker: {
                       image string
                       action enum # pull, start
                       ports [
                            {
                                guest int
                                host int
                            }
                       ]
                   }
                   // MUST NOT contains long running task in readiness check
                   ready: {
                        tasks []task
                   }
                   // TODO: health check
                }
            ]

        }
    ]
}
````

job status, a global view of job status, can also be used as final job history

- [ ] TODO: we need to know all stages in one node, and also need to know status of all nodes in one stage

````text
{
    id string
    rawSpec string
    spec JobSpec
    
    createTime int64  unix nano
    startTime int64
    stopTime int64
    
    status enum;
    currentPipeline # the array index in pipeline
    foregroundStages # array of index of stages currently running in foreground
    backgroundStages # array of index of stages currently running in background
    finishedStages # array of index of finished stages
    
    # TODO: how to describe pipelines
    pipelines: [
        {
            status enum # queued, running, finished, aborted, background
            stages: [
                # TODO: just refer index ?
            ]
        }
    ]
    stages [
        {
            index int # the index in stages, filled by central, not configurable by client
            pipeline int # the index in pipelines, filled by central, not configurable by client
            name string
            spec StageSpec # rendered with information
            nodes: [
                {
                    id string
                    name string # name assigned in config
                    role enum # role specified when got assigned
                    info NodeInfo # node info when registered to server
                    pipelines: [
                        {
                            status enum # queued, running, finished, aborted, background
                            tasks: [
                                # TODO: just refer index ?
                            ]
                        }
                    ]
                    tasks: [
                        foreground: [
                            {
                                startTime int64
                                # TODO: log, resource usage etc.
                            }
                        ]
                        background: [
                            {
                                startTime int64
                                readyTime int64
                            }
                        ]
                    ]
                }
            ]
            tasks: [
                {
                    nodes: [
                        {
                            status: finished?
                        }
                    ]
                }
            ]
        }
    ]
}
````