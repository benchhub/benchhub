# Host Metrics

TODO

- CPU, Memory on Host, per process metrics
- Network, Disk
- use fsnotify to monitor events in certain directories?
- stat
- iostat
- vmstat
- util for dealing with proc file system
  - i.e. an interface? file, then parse etc.
- shared memory `/dev/shm` https://gerardnico.com/wiki/linux/shared_memory
- https://osquery.io/schema/
- [ ] description of each metrics?

## Known issues

- machine disk space only consider filesystem mount to `/`

## Test data

````bash
cat /proc/stat > stat
````