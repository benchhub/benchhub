# Job

A job describes requirements of a benchmark and how it should be run

- spec https://github.com/benchhub/benchhub/issues/12
- [ ] TODO: multi file? allow specify framework specific stuff?
- [ ] might just use go file to create task config, install go environment on server

````yaml
owner: at15
project: oltpbench-test
framework: oltpbench
nodespec:
  - type: loader
    requirements:
      cpu:
        cores: 1
      mem:
        size: 1024
      disk:
        size: 20
      services:
        - name: docker
          version: 17.12
    count: 1
  - type: database
    requirements:
      cpu:
        cores: 1
      mem:
        size: 1024
      disk:
        size: 20
      services:
        - name: docker
          version: 17.12
     count: 1
nodes:
  - type: loader
    envs:
      - DB=mysql
      - BENCH=tpcc
  - type: loader
    envs:
      - DB=mysql
      - BENCH=tpcc
  - type: database
    ???  
````

pseudo config, we would be using go file to get it type check ...

- [ ] how to register config? generate a go file and run it? sounds good ...

````text
loader {
    stage prepare {
        shell {
            command bhagent install jdk
        }
    }
    stage ping {
        type: shell
        command: ab ping --remote=${central}
    }
    stage load {
        type: shell
        command: ab load --remote=${central}
    }
    stage teardown {
        // nothing
    }
}

database {
    stage setup {
        shell {
            command bhutil install jdk
        }
        shell {
            command bhutil ?
            failable true
        }
        shell {
            name xephonserve
            command xephonk serve
            type long running
        }
    }
    stage teardown {
        shell {
            kill xephonserve
        }
        shell {
            command rm -rf /var/lib/xephonk/data
        }
    }
}
````

generic framework

- setup
  - loader install software (i.e. jdk)
  - database start
- ping
  - all loader ping database
  - one loader do the initial loading if needed
  - after ack from all loader
- run
  - loader run until finish
- teardown

TODO: some database like KairosDB relies on Cassandra

task is like https://www.nomadproject.io/docs/drivers/docker.html

a job is made of several stages, some can be run in parallel, some should keep running until the end

````text
stage {
    name pull_docker_image
    name build_image ?
    // ...
}

stage {
    name start cassandra
    type long running
    node type=database
    command {
        driver docker
        image cassandra:3.1
        ports:
            guest:9042
            host:9042
    }
    ready {
        // TODO: someway to test cassandra is ready, i.e. wait for it https://github.com/benchhub/benchhub/issues/20
        // https://github.com/jwilder/dockerize
    }
    shutdown {
        docker stop cassandra
    }
}

stage {
    name start kairosdb
    type long running
    node type=database
    command {
        driver docker
        image ?
    }
    shutdown {
        docker stop kairosdb
    }
}

stage {
    name ping
    type short
    node type=worker
    command {
        libtsdb ping kairosdb --host=${central:port}
    }
}

stage {
    name load
    type short
    node type=worker
}

pipeline {
    pull_docker_image, 
    start_cassandra
    start_kairosdb
    ping
    load
    ....
}
````