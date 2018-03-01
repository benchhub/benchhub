# Mesos

- http://mesos.apache.org/documentation/latest/architecture/
- http://mesos.apache.org/documentation/latest/endpoints/
- http://mesos.apache.org/documentation/latest/app-framework-development-guide/

## Chronos

- https://mesos.github.io/chronos/docs/api.html#adding-a-docker-job

````json
{
  "schedule": "R/2014-09-25T17:22:00Z/PT2M",
  "name": "dockerjob",
  "container": {
    "type": "DOCKER",
    "image": "libmesos/ubuntu",
    "network": "BRIDGE",
    "volumes": [
      {
        "containerPath": "/var/log/",
        "hostPath": "/logs/",
        "mode": "RW"
      }
    ]
  },
  "cpus": "0.5",
  "mem": "512",
  "fetch": [],
  "command": "while sleep 10; do date =u %T; done"
}
````
