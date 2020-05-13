module github.com/benchhub/benchhub

go 1.14

require (
	github.com/dyweb/go.ice v0.0.3
	github.com/dyweb/gommon v0.0.13
	github.com/gogo/protobuf v1.3.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/xephonhq/xephon-b v0.3.0
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.2.7
)

replace github.com/dyweb/go.ice => ../../dyweb/go.ice
