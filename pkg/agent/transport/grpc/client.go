package grpc

import "google.golang.org/grpc"

func NewClient(con *grpc.ClientConn) BenchHubAgentClient {
	return NewBenchHubAgentClient(con)
}
