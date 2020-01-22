# Wait for a server is ready

- https://github.com/benchhub/benchhub/issues/20
- https://github.com/jwilder/dockerize

````text
Usage of waitforit:
  -h	display help
  -host value
    	hosts to wait
  -r int
    	number of retries for each host, negative for infinite retry (default 5)
  -t duration
    	time out for each connect (default 10s)
  -tt duration
    	total time out (default 1m0s)
  -w value
    	(alias of host) hosts to wait

Example:
  - Wait for two server using default config, return 0 when both of them can be connected
    waitforit -w tcp://localhost:9042 -w http://localhost:8080
````

TODO

- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/
- https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes