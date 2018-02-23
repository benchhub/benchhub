# Monitor

## Take away

- use `strace -e trace=file df` to see syscall on file

## Related Issues

- [Xephon-K#39 Collector](https://github.com/xephonhq/xephon-k/issues/39)
  - https://github.com/xephonhq/xephon-k/projects/2

## Survey

- [cAdvisor](cadvisor.md)

## What we want

- general metrics, CPU, Network, Disk usage
- process specific, i.e. how many bytes written (to observe write amplification, the final file size does not count)