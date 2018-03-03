# Wait for a server is ready

- https://github.com/benchhub/benchhub/issues/20
- https://github.com/jwilder/dockerize

````go
func waitForSocket(scheme, addr string, timeout time.Duration) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			conn, err := net.DialTimeout(scheme, addr, waitTimeoutFlag)
			if err != nil {
				log.Printf("Problem with dial: %v. Sleeping %s\n", err.Error(), waitRetryInterval)
				time.Sleep(waitRetryInterval)
			}
			if conn != nil {
				log.Printf("Connected to %s://%s\n", scheme, addr)
				return
			}
		}
	}()
}
````