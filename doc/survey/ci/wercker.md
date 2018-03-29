# Wercker

- bought by oracle ...
- support dev (watch and reload)
- each script has `name` and `code` instead of just one line, verbose, but it is what I wanted when building Ayi

````yaml
# The container definition we want to use for developing our app
box:
  id: golang
  ports:
    - "5000"

# Defining the dev pipeline
dev:
  steps:
    - internal/watch:
        code: |
          go build ./...
          ./source
        reload: true

build:
  steps:
    - wercker/golint
    - script:
        name: go build
        code: |
          go build ./...
    - script:
        name: go test
        code: |
          go test ./...
````