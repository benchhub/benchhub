package grpc

import "google.golang.org/grpc"

func NewClient(con *grpc.ClientConn) BenchHubCentralClient {
	return NewBenchHubCentralClient(con)
}
