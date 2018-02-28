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