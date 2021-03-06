# Alibaba

- Fuxi https://github.com/gaocegege/papers-notebook/issues/19
- https://edu.aliyun.com/course/31
  - by @gaocegege '阿里的分布式调度课程，讲的说实话很好了'
  - client -> fuxi master -> app master -> app worker
  - app master 
    - require resource from fuxi master
    - start app worker on node
    - monitor job, retry
  - app worker
    - just run and store result
  - job scheduling
    - run job, monitor, restart etc.
    - locality?
  - resource scheduling
    - priority
    - 抢占 ...
    - fair scheduling, multiple groups based on priority
    - quota, group jobs, fuxi would change quota dynamically
       - make sharable resource cheaper
       - make dedicated resource more expensive
       - **it's a bill management system** 
  - fault tolerance
    - job scheduling
      - app worker, retry on another app worker
    - resource scheduling
      - soft state (recover from other components)
        - ask app master and tubo to send snapshot of state
      - hard state (i.e. job config)
        - check point
  - scale
    - multi thread, on app master, use different thread pools, one for fuximaster (more import & less node), one for tubo
    - incremental, (keep memory of app master in fuxi master)
  - security
    - static: encrypt messages
    - dynamic: token from authority
    - sandbox in tubo when start worker
  - isolation
    - vm (KVM), best isolation
    - docker
    - lxc
  - future
    - mix of online and offline
    - real-time compute, storm, spark
    - larger scale
      - time
      - resource