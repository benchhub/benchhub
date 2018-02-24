package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HandleFunc func(ctx context.Context, req interface{}) (res interface{}, err error)

type Handler interface {
	EmptyRequest() interface{}
	Handle() HandleFunc
}

var _ Handler = (*handler)(nil)

type handler struct {
	f func() interface{}
	h HandleFunc
}

func (h *handler) EmptyRequest() interface{} {
	return h.f()
}

func (h *handler) Handle() HandleFunc {
	return h.h
}

type HandlerRegister func(mux *Mux)

type Mux struct {
	handlers map[string]Handler
}

func (m *Mux) AddHandler(path string, factory func() interface{}, h HandleFunc) {
	m.handlers[path] = &handler{
		f: factory,
		h: h,
	}
}

func NewHttpHandler(register HandlerRegister) http.Handler {
	mux := &Mux{
		handlers: make(map[string]Handler, 10),
	}
	register(mux)
	httpMux := http.NewServeMux()
	for path, h := range mux.handlers {
		httpMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			req := h.EmptyRequest()
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				invalidRequestFormat(w, err)
				return
			}
			handle := h.Handle()
			res, err := handle(ctx, req)
			if err != nil {
				internalError(w, err)
				return
			}
			buf := &bytes.Buffer{}
			if err := json.NewEncoder(buf).Encode(res); err != nil {
				internalError(w, err)
				return
			}
			if _, err := w.Write(buf.Bytes()); err != nil {
				log.Warnf("can't write to http connection %v", err)
			}
		})
	}
	return httpMux
}

func invalidRequestFormat(w http.ResponseWriter, err error) {
	w.WriteHeader(400)
	// TODO: this should be json, if request is json
	fmt.Fprintf(w, "invalid request json: %v", err)
}

func internalError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	// TODO: this should be json, if request is json
	fmt.Fprintf(w, "internal erro :%v", err)
}
