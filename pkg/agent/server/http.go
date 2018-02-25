package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dyweb/gommon/errors"

	ihttp "github.com/at15/go.ice/ice/transport/http"
	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"
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

func (srv *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	jMux := ihttp.NewJsonHandlerMux()
	srv.RegisterHandler(jMux)
	jMux.MountToStd(mux)
	return mux
}

func (srv *HttpServer) RegisterHandler(mux *ihttp.JsonHandlerMux) {
	mux.AddHandlerFunc("/ping", func() interface{} {
		return &pbc.Ping{}
	}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		if ping, ok := req.(*pbc.Ping); !ok {
			return nil, errors.New("invalid type, can't cast to *pbc.Ping")
		} else {
			return srv.Ping(ctx, ping)
		}
	})
}
