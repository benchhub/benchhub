# Host Metrics

TODO

- CPU, Memory on Host, per process metrics
- use fsnotify to monitor events in certain directories?
- stat
- iostat
- vmstat
- util for dealing with proc file system
  - i.e. an interface? file, then parse etc.

## Test data

````bash
cat /proc/stat > stat
````