# Apache Aurora

- http://aurora.apache.org/documentation/latest/getting-started/overview/

## Overview

- aurora manages jobs made of tasks
- mesos manage tasks made of processes
- [ ] thermos manages processes?
- defined in `.aurora` configuration file using https://github.com/wickman/pystachio

## Tutorial

- http://aurora.apache.org/documentation/latest/getting-started/tutorial/

````python
pkg_path = '/vagrant/hello_world.py'

# we use a trick here to make the configuration change with
# the contents of the file, for simplicity.  in a normal setting, packages would be
# versioned, and the version number would be changed in the configuration.
import hashlib
with open(pkg_path, 'rb') as f:
  pkg_checksum = hashlib.md5(f.read()).hexdigest()

# copy hello_world.py into the local sandbox
install = Process(
  name = 'fetch_package',
  cmdline = 'cp %s . && echo %s && chmod +x hello_world.py' % (pkg_path, pkg_checksum))

# run the script
hello_world = Process(
  name = 'hello_world',
  cmdline = 'python -u hello_world.py')

# describe the task
hello_world_task = SequentialTask(
  processes = [install, hello_world],
  resources = Resources(cpu = 1, ram = 1*MB, disk=8*MB))

jobs = [
  Service(cluster = 'devcluster',
          environment = 'devel',
          role = 'www-data',
          name = 'hello_world',
          task = hello_world_task)
]
````

## Configuration

http://aurora.apache.org/documentation/latest/reference/configuration-tutorial/

- process define a single command
- task has multiple processes, can be run in parallel (default) or in order (using constraint)
  - `Tasks.combine` make sure one success before run another
  - `Tasks.concat` run in parallel
  - resources
    - cpu number of cores
    - ram bytes
    - disk bytes
- job can only take one task
  - role
  - environment
  - cluster
