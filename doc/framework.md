# Framework

Framework is used to solve the dependency problem across nodes, i.e. workload generator should only start after the database
is already started, if database crashed, workload generator should be stopped as well

The simplest example, start a http server before the start running workload

- in parallel, start workload generator (if it need to load data into memory etc.) and server
- let all loader ping server (somehow), report their success to framework
- framework tell all loader to start benchmark at a given time, i.e. 1 min later, all loader send ack, framework tell all loader all loader is ready, otherwise loader should abort

generic framework

- run shell commands and container?

stages

- download binary
- start database
- warm up, download dataset
- make sure client can connect to database
- run small workload
- run real workload

... hmm ... still kind of mixing stuff together

loader {
  starts_when database.ready == true
}

database {
  stop_when loader.finished == true
}

- prepare
  - loader