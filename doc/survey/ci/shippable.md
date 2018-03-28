# Shippable

- https://www.shippable.com/
- matrix is similar to travis

````yaml
language:

node_js:      # language runtime
  - #language version

services:
  - #any supported service

env:
  - #env1=foo
  - #env2=bar

matrix:

build:

  pre_ci:

  pre_ci_boot:
    image_name:
    image_tag:
    pull:
    options:
  ci:
    - npm install    
    - mkdir -p shippable/testresults
    - mkdir -p shippable/codecoverage
    - mysql -e 'create database if not exists test;'
    - grunt
    - npm test
  post_ci:
    - #command1
    - #command2
  on_success:
    - #command1
    - #command2
  on_failure:
    - #command1
    - #command2
  cache:
  cache_dir_list:
    - #dir1
  push:
    - #command1

integrations:
 notifications:
   - integrationName:
     type:
     recipients:
       - #recp1
       - #recp2

  hub:
    - integrationName:
      type:
      agent_only:
````