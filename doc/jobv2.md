# Job v2

- job is a benchmark job including multiple stages
  - ~~provision required nodes using IaaS~~
  - acquire nodes from scheduler 
  - install the database (pull the docker image if using docker)
  - install workload generator
  - start the database
  - create required schema
  - make sure loader can reach database
  - run workload
  - clean up
  - a final report
- each stage can contains several tasks
  - two type of tasks, run and finish (install, load), and long running (database)
  - run and finish should all success otherwise the stage fails
  - long running should be stopped later
  
A complex example of benchmark KairosDB using Xephon-B, it is complex because

- KairosDB requires Cassandra up and running, otherwise it fails directly when start (it might have retry, but Cassandra must be reachable when start)
- Xephon-B can use multiple nodes as loader, we need to make sure they all can reach database and start at same time

TODO

- [ ] how to assign nodes when job start

````text
name xephon-b-v0.1-kairosdb-1.2.0
reason triggered by ci
workload xephon-b
database kairosdb

stage {
    name install Xephon-B
    nodes {
        type = loader
    }
    tasks {
        task {
            driver shell
            curl github.com/xephonhq/xephon-b/v0.0.1.tar.gz && tar zxvf ... && mv xb /usr/bin
        }
        task {
            driver shell
            curl xb version
        }
    }
}

stage {
    name install Cassandra
    nodes {
        type = database
    }
    tasks {
        task {
            driver shell
            docker pull cassandra:3.1
        }
        task {
            driver shell
            docker image list
        }
    }
}

stage {
    name build KairosDB
    nodes {
        type = database
    }
    tasks {
        task {
            driver shell
            git clone github.com/benchhub/cookbook
        }
        task {
            driver shell
            docker build -t benchhub/kairosdb:lastest .
        }
        task {
            docker image list
        }
    }
}

stage {
    name start Cassandra
    nodes {
        type = database
    }
    type long
    tasks {
        task {
            driver docker
            docker start cassandra:3.1 cassandra
            type long 
            ready {
                checkhostport localhost:9042 
            }
            shutdown {
                docker stop cassandra
            }
        }
    }
}

stage {
    name start KairosDB
    nodes {
        type = database
    }
    type long 
    tasks {
        task {
            driver docker
            docker start benchhub/kairosdb:lastest kairosdb
            type long
            ready {
                checkhostport localhost:8080
            }
            shutdown {
                docker stop kairosdb
            }
        }
        
    }
}

stage {
    name all ping
    nodes {
        type = loader
    }
    tasks {
        task {
            driver shell
            libtsdb ping ${central} --db=kairosdb
            retry 3
        }
    }
}

stage {
    name load
    nodes {
        type = loader
    }
    tasks {
        task {
            driver shell
            xb --target=${central} --type=kairosdb
        }
    }
}

pipeline = [
    [install Xephon-B, install Cassandra, install KairosDB],
    [start Cassandra],
    [start KairosDB],
    [load]
]
````

