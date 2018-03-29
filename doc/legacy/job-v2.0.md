# Job

- version 2
- created on 2018/02/28 by @at15
- updated on 2018/03/01 by @at15

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
  - long running (background, daemon) should be stopped later
  
A complex example of benchmark KairosDB using Xephon-B, it is complex because

- KairosDB requires Cassandra up and running, otherwise it fails directly when start (it might have retry, but Cassandra must be reachable when start)
- Xephon-B can use multiple nodes as loader, we need to make sure they all can reach database and start at same time

TODO

- [x] how to assign nodes when job start, specify resources and add node selector for stages
  - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

NOTE: the real spec is currently written in YAML, see [pkg/common/spec/xephonb-kairosdb.yml](../pkg/common/spec/xephonb-kairosdb.yml)
NOTE: an updated version is in [store.md](store.md)

````text
name xephon-b-v0.1-kairosdb-1.2.0
reason triggered by ci
framework xephon-b
database kairosdb

nodes {
    node {
        name db
        type database
        resource {}
    }
    node {
        name loader1
        type loader
        resource {}
    }
    node {
        name loader2
        type loader
        resource {}
    }
}

stage {
    name install Xephon-B
    selectors [
        { type = loader }
    ]
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
    selectors [
        { type = database }
    ]
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
    selectors [
        { type = database }
    ]
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
    selectors [
        { type = database }
    ]
    background true
    tasks {
        task {
            driver docker
            docker start cassandra:3.1 cassandra
            background true 
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
    selectors {
        type = database
    }
    background true 
    tasks {
        task {
            driver docker
            docker start benchhub/kairosdb:lastest kairosdb
            background true
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
    selectors [
        { type = loader }
    ]
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
    selectors [
        { type = loader }
    ]
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
    [all ping],
    [load]
]
````

