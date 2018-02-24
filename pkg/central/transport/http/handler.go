package http

import (
	"context"
	"net/http"
)

type HandlerRegister func(mux *Mux)

type Mux struct {
	handlers map[string]HandlerE
}

type HandlerE func(ctx context.Context, req interface{}) (res interface{}, err error)

func NewHttpHandler(register HandlerRegister) http.Handler {
	mux := &Mux{
		handlers: make(map[string]HandlerE, 10),
	}
	register(mux)
	httpMux := http.NewServeMux()
	for path, h := range mux.handlers {
		httpMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			// TODO: a decoder is required ... an encoder is not ...
			// we are going back to the old ways of go-kit
		})
	}
	return mux
}
