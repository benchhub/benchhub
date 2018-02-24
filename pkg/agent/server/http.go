package server

import (
	"context"
	"os"
	"fmt"

	"github.com/pkg/errors"

	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
	"net/http"
)

// HttpServer is mainly used to communicate with browser, routes are mounted in transport http package
type HttpServer struct {
}

func NewHttpServer() (*HttpServer, error) {
	return &HttpServer{}, nil
}

func (srv *HttpServer) Ping(ctx context.Context, ping *pbc.Ping) (*pbc.Pong, error) {
	if host, err := os.Hostname(); err != nil {
		return &pbc.Pong{Message: "pong from unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		res := fmt.Sprintf("pong from agent %s your message is %s", host, ping.Message)
		return &pbc.Pong{Message: res}, nil
	}
}

func (srv *HttpServer) HandlerRegister(mux *http.ServeMux) {
	mux.Handle("/ping", http.)
}