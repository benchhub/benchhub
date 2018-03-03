package waitforit

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/dyweb/gommon/errors"
)

func WaitSockets(ctx context.Context, us []url.URL, totalTimeout time.Duration, timeout time.Duration, retry int) (time.Duration, error) {
	start := time.Now()
	if retry < 0 {
		retry = 66666 // a magic big number
	}
	var wg sync.WaitGroup
	wg.Add(len(us))
	merr := errors.NewMultiErrSafe()
	ctx, cancel := context.WithTimeout(ctx, totalTimeout)
	for _, u := range us {
		go func(u url.URL) {
			i := 0
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					if i < retry {
						log.Printf("retry %s %d", u.String(), i)
						_, err := WaitSocket(u, timeout)
						if err == nil {
							log.Printf("%s is up", u.String())
							wg.Done()
							return
						} else {
							log.Printf("%s %v", u.String(), err)
							time.Sleep(timeout)
						}
						i++
					} else {
						merr.Append(errors.Errorf("%s exceed max retry %d", u.String(), retry))
						wg.Done()
						cancel()
						return
					}
				}
			}
		}(u)
	}
	wg.Wait()
	return time.Now().Sub(start), merr.ErrorOrNil()
}

func WaitSocket(u url.URL, timeout time.Duration) (time.Duration, error) {
	log.Printf("wait for socket %s %s", u.String(), timeout)
	switch u.Scheme {
	case "tcp", "tcp4", "tcp6":
		return waitRaw(u.Scheme, u.Host, timeout)
	case "unix":
		return waitRaw(u.Scheme, u.Path, timeout)
	case "http", "https":
		return waitHttp(u, timeout)
	}
	return timeout, errors.Errorf("unsupported schema %s", u.Scheme)
}

func waitRaw(schema string, addr string, timeout time.Duration) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout(schema, addr, timeout)
	if err != nil {
		return timeout, err
	}
	defer conn.Close()
	if conn != nil {
		return time.Now().Sub(start), nil
	}
	return timeout, errors.New("nil connection with nil error")
}

// TODO: allow insecure
// TODO: allow certain status code, sometimes we just want to make sure the server is up
func waitHttp(u url.URL, timeout time.Duration) (time.Duration, error) {
	start := time.Now()
	c := http.Client{
		Transport: NewDefaultTransport(),
		Timeout:   timeout,
	}
	res, err := c.Get(u.String())
	if err != nil {
		return timeout, err
	}
	// TODO: read body, check if valid json etc.
	//b, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	if res.StatusCode >= 200 && res.StatusCode <= 300 {
		return time.Now().Sub(start), nil
	}
	return timeout, errors.Errorf("invalid status code %d", res.StatusCode)
}

// Default Transport Client that is same as https://golang.org/src/net/http/transport.go
// It's similar to https://github.com/hashicorp/go-cleanhttp

// NewDefaultTransport is copied from net/http/transport.go
func NewDefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
