# Xephon-K

Xephon-K also has its own collector, though I think it's never put into real use because its UI is never finished

- https://github.com/xephonhq/xephon-k/issues/39
- https://github.com/xephonhq/xephon-k/issues/21 contains some reference
  - https://github.com/prometheus/node_exporter/blob/master/collector/stat_linux.go#L87
  - https://github.com/shirou/gopsutil is using `/proc` if I remembered correctly
    - used by https://github.com/influxdata/telegraf/blob/master/plugins/inputs/system/cpu.go 
  - https://github.com/cloudfoundry/gosigar
  - https://github.com/elastic/beats/tree/master/metricbeat/module/system
  - GoSigar fork w/ cgroup by Elastic https://github.com/elastic/gosigar/tree/master/cgroup 