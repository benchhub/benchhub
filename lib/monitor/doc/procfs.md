# The proc filesystem

## Reference

- https://www.kernel.org/doc/Documentation/filesystems/proc.txt
- http://manpages.ubuntu.com/manpages/zesty/en/man5/proc.5.html

## Static

Hardware information and other values that don't change after boot

## General

### /proc/stat

- `cpu`: user, nice, system, idle, iowait, irq, softirq, steal, quest, quest_nice
- `intr`: counts of interrupts, .... (pretty long, cut in example)
- `ctxt`: context switches
- `btime`: boot time in seconds since Epoch
- `process`: number of fork since boot
- `procs_running`: number of process in runnable state
- `procs_blocked`: number of process blocked waiting for I/O to complete
- `page`, `swap`, `disk_io` are not found on Ubuntu 17.04 4.10
- `softirq` is not found in man page

````text
cpu  79894 810 21788 1689252 4158 0 1184 0 0 0
cpu0 10176 35 2365 211402 451 0 436 0 0 0
cpu1 10226 31 2411 211672 351 0 201 0 0 0
cpu2 10116 47 2105 212068 286 0 84 0 0 0
cpu3 10223 106 2164 211676 227 0 162 0 0 0
cpu4 10575 124 2362 211388 248 0 13 0 0 0
cpu5 10244 45 2234 211807 264 0 8 0 0 0
cpu6 10698 347 2904 208660 1886 0 268 0 0 0
cpu7 7633 72 5240 210575 442 0 9 0 0 0
intr 2636725 23 3 0 0 0 0 0 0 1 0 0 0 4 0 0 0 0 1930 0 0 0 0 0
ctxt 10447493
btime 1514922734
processes 14411
procs_running 1
procs_blocked 0
softirq 2379472 2 646290 436 210842 305639 0 642 681892 0 533729
````