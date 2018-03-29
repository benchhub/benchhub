# Travis

- https://github.com/travis-ci/travis-hub
- https://github.com/travis-ci/travis-scheduler
- https://github.com/travis-ci/worker

An example from our fork of OLTPBench

- env can be used to form a build matrix
  - when run locally, in case of benchmark, they (cells in matrix) can only be run in serial
- several stages
- seems commands are executed using `sh -c` or generated sh 
- `cd` has permanent effect, need to keep an eye on that
 
````yaml
git:
  depth: 3
sudo: required
services:
  - docker
language: java # default is oracle JDK 1.8.0_151
jdk:
  - oraclejdk8
addons:
  apt:
    packages:
      - "python3"
      - "python3-pip"
env:
# --- begin tpcc with different databases ---
  - BENCH=tpcc DB=mysql
  - BENCH=tpcc DB=postgres
  - BENCH=tpcc DB=tidb
# https://stackoverflow.com/questions/34377017/what-are-the-differences-between-the-before-install-script-travis-yml-opti
before_install:
  - python3 --version
  - sudo pip3 install pyyaml
  - java -version
  - docker version
  - docker-compose version
  # check if common ports are already in use, travis starts a bunch of DBMS by default ....
  - sudo netstat -nlp | grep :3306
  - sudo netstat -nlp | grep :5432
  - sudo service mysql stop
  - sudo service postgresql stop
# https://docs.travis-ci.com/user/database-setup/#Multiple-Database-Builds
install:
  - ant build
  - ./docker/travis_start.sh
before_script:
  - ./config/config.py validate
script:
  - ./config/config.py generate --bench ${BENCH} --db ${DB} --scalefactor 1
  - ./oltpbenchmark --bench ${BENCH} --config config/generated_${BENCH}_${DB}_config.xml --create true --load true --execute true
after_script:
  - ./docker/travis_stop.sh
````