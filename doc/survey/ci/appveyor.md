# AppVeyor

- https://www.appveyor.com/
  - complete yaml https://www.appveyor.com/docs/appveyor-yml/
- build matrix support **Build cloud**, **Platform**, **OS**
- a build, deploy pipeline, kind of like circleci

````yaml
environment:
  # these variables are common to all jobs
  common_var1: value1
  common_var2: value2
  matrix:
    # first group
    - db: mysql
      provider: mysql

    # second group
    - db: mssql
      provider: mssql
      password:
        secure: DHEU39J6X9VD376==
platform:
  - x86
  - Any CPU
configuration:
  - Debug
  - Release
matrix:
  fast_finish: true
````