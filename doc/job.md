# Job

A job describes requirements of a benchmark and how it should be run

- spec https://github.com/benchhub/benchhub/issues/12
- [ ] TODO: multi file? allow specify framework specific stuff?

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