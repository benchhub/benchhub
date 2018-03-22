# Job

## Spec

````text
{
    id string # assigned when job is created
    name string 
    owner: {
        id string
        name string
        type enum # user, org
    }
    workload: {
        framework string
        frameworkVersion string
        database string
        databaseVersion string        
    }
    nodeAssignments: [
        {
            name string # the unique name used throughout config TODO: might put it into dns?
            role enum # node role, loader or database
            # TODO: add label to nodes
            # TODO: select node etc. by resource etc.
            selectors [
            ]
        }
    ]
    pipelines: [
        {
            name string # unique in pipelines
            stages []string # stages that can run in parallel
        }
    ]
    stages: [
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

## Scheduler (Node Assignment)

- currently just using role

TODO

- use node state
- use node label

After node assignment, result is returned in the order of spec

## Planner 

- a serial of execution

````text
[
    {
        name: pipelinename
        stages: [
            {
                nodes: [{ip: xxx, name: xxx} ... ]
                pipelines: [
                    {
                      name:
                      tasks: [
                      
                      ]
                    },
                 ]
            },
            {
            
            }
        ]
    },
    {
    
    }
]
````

Copied from planner_test using json format

````json
{
   "pipelines":[
      {
         "name":"download",
         "stages":[
            {
               "nodes":[
                  {
                     "node":{
                        "id":"id-loader",
                        "role":3,
                        "state":2,
                        "info":{
                           "id":"id-loader",
                           "addr":{

                           },
                           "role":3,
                           "provider":{

                           },
                           "capacity":{

                           }
                        }
                     },
                     "spec":{
                        "name":"cli",
                        "role":3,
                        "selectors":null
                     }
                  }
               ],
               "pipelines":[
                  {
                     "name":"autogen-0",
                     "tasks":[
                        {
                           "spec":{
                              "driver":2,
                              "shell":{
                                 "command":"wget https://github.com/benchhub/benchhub/releases/download/v0.0.1/pingclient-0.0.1.zip \u0026\u0026 unzip pingclient-0.0.1.zip"
                              },
                              "ready":{
                                 "tasks":null
                              }
                           }
                        }
                     ]
                  }
               ]
            },
            {
               "nodes":[
                  {
                     "node":{
                        "id":"id-database",
                        "role":4,
                        "state":2,
                        "info":{
                           "id":"id-database",
                           "addr":{

                           },
                           "role":4,
                           "provider":{

                           },
                           "capacity":{

                           }
                        }
                     },
                     "spec":{
                        "name":"srv",
                        "role":4,
                        "selectors":null
                     }
                  }
               ],
               "pipelines":[
                  {
                     "name":"autogen-0",
                     "tasks":[
                        {
                           "spec":{
                              "driver":2,
                              "shell":{
                                 "command":"wget https://github.com/benchhub/benchhub/releases/download/v0.0.1/pingserver-0.0.1.zip \u0026\u0026 unzip pingserver-0.0.1.zip"
                              },
                              "ready":{
                                 "tasks":null
                              }
                           }
                        }
                     ]
                  }
               ]
            }
         ]
      },
      {
         "name":"run_server",
         "stages":[
            {
               "nodes":[
                  {
                     "node":{
                        "id":"id-database",
                        "role":4,
                        "state":2,
                        "info":{
                           "id":"id-database",
                           "addr":{

                           },
                           "role":4,
                           "provider":{

                           },
                           "capacity":{

                           }
                        }
                     },
                     "spec":{
                        "name":"srv",
                        "role":4,
                        "selectors":null
                     }
                  }
               ],
               "pipelines":[
                  {
                     "name":"autogen-0",
                     "tasks":[
                        {
                           "spec":{
                              "background":true,
                              "driver":2,
                              "shell":{
                                 "command":"pingserver 8080"
                              },
                              "ready":{
                                 "tasks":[
                                    {
                                       "driver":2,
                                       "shell":{
                                          "command":"waitforit -w http://localhost:8080/ping"
                                       },
                                       "ready":{
                                          "tasks":null
                                       }
                                    }
                                 ]
                              }
                           }
                        }
                     ]
                  }
               ]
            }
         ]
      },
      {
         "name":"run_workload",
         "stages":[
            {
               "nodes":[
                  {
                     "node":{
                        "id":"id-loader",
                        "role":3,
                        "state":2,
                        "info":{
                           "id":"id-loader",
                           "addr":{

                           },
                           "role":3,
                           "provider":{

                           },
                           "capacity":{

                           }
                        }
                     },
                     "spec":{
                        "name":"cli",
                        "role":3,
                        "selectors":null
                     }
                  }
               ],
               "pipelines":[
                  {
                     "name":"autogen-0",
                     "tasks":[
                        {
                           "spec":{
                              "driver":2,
                              "shell":{
                                 "command":"pingclient http://{{.Nodes.srv.Ip}}:8080"
                              },
                              "ready":{
                                 "tasks":null
                              }
                           }
                        }
                     ]
                  }
               ]
            }
         ]
      }
   ]
}
````

## Executor

Both central and agent has job executor, the one in central is a dispatcher and monitor, the one in agent does the actual execution

Central

- requires generated plan
- [ ] how to dispatch job to mock executor?
- pick the first stage
- for all the selected nodes in this stage
- send the plan of this stage to each nodes
- when central get update about job, it is dispatched to correspond job executor
  - [x] the server registry keeps a map of job manager
  
Agent

- listen to plan sent from server
- make sure one job only have on running stage
  - [ ] return error code for Conflict?
