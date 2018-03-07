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
    NodeAssignment: [
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