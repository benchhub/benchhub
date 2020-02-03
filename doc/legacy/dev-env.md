# Develop environment setup

## Requirement

- OS
  - Linux/Mac, Windows is not supported due to monitoring lib is using `/proc` to collect metrics
    - use vagrant should work, but [script/vagrant](../script/vagrant) is for a local 3 node setup using compiled binaries, it's not for dev on windows
- Go 
  - version 1.9+
  - dep https://github.com/golang/dep `go get -u github.com/golang/dep/cmd/dep`
  - gommon https://github.com/dyweb/gommon `go get -u github.com/dyweb/gommon/cmd/gommon`
- Protobuf 
  - compiler (protoc)
  - gogoprotobuf in your $GOPATH
  
## IDE

- Goland
  - protobuf plugin
    - add `$GOPATH` to include path