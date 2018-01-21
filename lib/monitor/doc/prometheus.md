# Prometheus

- https://github.com/prometheus/client_golang
- https://github.com/prometheus/node_exporter
- https://prometheus.io/docs/concepts/metric_types/
  - https://prometheus.io/docs/practices/histograms/
- [Prometheus: Design and Philosophy - why it is the way it is](https://youtu.be/QgJbxCWRZ1s?t=29m31s)
  - also showed example of promql

Metric Types

> The Prometheus server does not yet make use of the type information and flattens all data into untyped time series. This may change in the future.

- counter: only goes up
- gauge: up and down
- histogram & summary
  - bucket and window?
  - [ ] TODO: calculate on client side? or server side?
  - https://github.com/prometheus/client_java/issues/328
  
````go
// Collector is the interface a collector has to implement.
type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update(ch chan<- prometheus.Metric) error
}
````

- `go run node_exporter.go` and visit http://localhost:9100/metrics, [example](prometheus_node_exporter.txt)

````go
// NodeCollector implements the prometheus.Collector interface.
type nodeCollector struct {
	Collectors map[string]Collector
}

// Collect implements the prometheus.Collector interface.
func (n nodeCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(n.Collectors))
	for name, c := range n.Collectors {
		go func(name string, c Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}

func execute(name string, c Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := c.Update(ch)
	duration := time.Since(begin)
	var success float64

	if err != nil {
		log.Errorf("ERROR: %s collector failed after %fs: %s", name, duration.Seconds(), err)
		success = 0
	} else {
		log.Debugf("OK: %s collector succeeded after %fs.", name, duration.Seconds())
		success = 1
	}
	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, duration.Seconds(), name)
	ch <- prometheus.MustNewConstMetric(scrapeSuccessDesc, prometheus.GaugeValue, success, name)
}
````

````go
type meminfoCollector struct{}

func init() {
	registerCollector("meminfo", defaultEnabled, NewMeminfoCollector)
}

// NewMeminfoCollector returns a new Collector exposing memory stats.
func NewMeminfoCollector() (Collector, error) {
	return &meminfoCollector{}, nil
}

// Update calls (*meminfoCollector).getMemInfo to get the platform specific
// memory metrics.
func (c *meminfoCollector) Update(ch chan<- prometheus.Metric) error {
	memInfo, err := c.getMemInfo()
	if err != nil {
		return fmt.Errorf("couldn't get meminfo: %s", err)
	}
	log.Debugf("Set node_mem: %#v", memInfo)
	for k, v := range memInfo {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				prometheus.BuildFQName(namespace, memInfoSubsystem, k),
				fmt.Sprintf("Memory information field %s.", k),
				nil, nil,
			),
			prometheus.GaugeValue, v,
		)
	}
	return nil
}

func (c *meminfoCollector) getMemInfo() (map[string]float64, error) {
	file, err := os.Open(procFilePath("meminfo"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseMemInfo(file)
}

func parseMemInfo(r io.Reader) (map[string]float64, error) {
	var (
		memInfo = map[string]float64{}
		scanner = bufio.NewScanner(r)
		re      = regexp.MustCompile(`\((.*)\)`)
	)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		fv, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value in meminfo: %s", err)
		}
		key := parts[0][:len(parts[0])-1] // remove trailing : from key
		// Active(anon) -> Active_anon
		key = re.ReplaceAllString(key, "_${1}")
		switch len(parts) {
		case 2: // no unit
		case 3: // has unit, we presume kB
			fv *= 1024
			key = key + "_bytes"
		default:
			return nil, fmt.Errorf("invalid line in meminfo: %s", line)
		}
		memInfo[key] = fv
	}

	return memInfo, scanner.Err()
}
````