# gopsutil

https://github.com/shirou/gopsutil which is a port of https://github.com/giampaolo/psutil

host/host_linux.go

- /etc/os-release
- /etc/lsb-release

cpu/cpu_linux.go

- /usr/bin/getconf CLK_TCK
- /proc/stat
- /proc/cpuinfo
- /sys/devices/system/cpu/cpu0

mem/mem_linux.go

- /proc/meminfo
- /proc/vmstat

disk/disk_linux.go

- /proc/mounts
- /proc/filesystems
- /proc/diskstats

net/net_linux.go

- /proc/net/dev
- /proc/net/snmp
- /proc/sys/net/netfilter/nf_conntrack_count
- /proc/sys/net/netfilter/nf_conntrack_max

load

- /proc/loadavg
- /proc/stat for procs_running, prcos_blocked, ctxt

process

- /proc/pid
- /proc/pid/fd
- /proc/pid/io