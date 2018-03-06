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