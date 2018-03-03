# Telegraf

- it use https://github.com/shirou/gopsutil for system metrics
- input interface, `Gather` is called periodically in `agent/agent.go` 

````go
type Input interface {
	// SampleConfig returns the default configuration of the Input
	SampleConfig() string

	// Description returns a one-sentence description on the Input
	Description() string

	// Gather takes in an accumulator and adds the metrics that the Input
	// gathers. This is called every "interval"
	Gather(Accumulator) error
}
````

````go
// inputs/system/memory.go
type SwapStats struct {
	ps PS
}

func (_ *SwapStats) Description() string {
	return "Read metrics about swap memory usage"
}

func (_ *SwapStats) SampleConfig() string { return "" }

func (s *SwapStats) Gather(acc telegraf.Accumulator) error {
	swap, err := s.ps.SwapStat()
	if err != nil {
		return fmt.Errorf("error getting swap memory info: %s", err)
	}

	fieldsG := map[string]interface{}{
		"total":        swap.Total,
		"used":         swap.Used,
		"free":         swap.Free,
		"used_percent": swap.UsedPercent,
	}
	fieldsC := map[string]interface{}{
		"in":  swap.Sin,
		"out": swap.Sout,
	}
	acc.AddGauge("swap", fieldsG, nil)
	acc.AddCounter("swap", fieldsC, nil)

	return nil
}

func init() {
	ps := newSystemPS()
	inputs.Add("mem", func() telegraf.Input {
		return &MemStats{ps: ps}
	})

	inputs.Add("swap", func() telegraf.Input {
		return &SwapStats{ps: ps}
	})
}
````

````go
// agent/agent.go

// Agent runs telegraf and collects data based on the given config
type Agent struct {
	Config *config.Config
}

// gatherWithTimeout gathers from the given input, with the given timeout.
//   when the given timeout is reached, gatherWithTimeout logs an error message
//   but continues waiting for it to return. This is to avoid leaving behind
//   hung processes, and to prevent re-calling the same hung process over and
//   over.
func gatherWithTimeout(
	shutdown chan struct{},
	input *models.RunningInput,
	acc *accumulator,
	timeout time.Duration,
) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	done := make(chan error)
	go func() {
		done <- input.Input.Gather(acc)
	}()

	for {
		select {
		case err := <-done:
			if err != nil {
				acc.AddError(err)
			}
			return
		case <-ticker.C:
			err := fmt.Errorf("took longer to collect than collection interval (%s)",
				timeout)
			acc.AddError(err)
			continue
		case <-shutdown:
			return
		}
	}
}
````

## Docker

inputs/docker/docker.go  does not use stream

````go
func (d *Docker) gatherContainer(container types.Container,	acc telegraf.Accumulator) error {
	var v *types.StatsJSON
	tags := map[string]string{
  		"engine_host":       d.engine_host,
  		"container_name":    cname,
  		"container_image":   imageName,
  		"container_version": imageVersion,
  	}
  
  	if !d.containerFilter.Match(cname) {
  		return nil
  	}
  
  	ctx, cancel := context.WithTimeout(context.Background(), d.Timeout.Duration)
  	defer cancel()
  	r, err := d.client.ContainerStats(ctx, container.ID, false)
  	if err != nil {
  		return fmt.Errorf("Error getting docker stats: %s", err.Error())
  	}
  	defer r.Body.Close()
  	dec := json.NewDecoder(r.Body)
  	if err = dec.Decode(&v); err != nil {
  		if err == io.EOF {
  			return nil
  		}
  		return fmt.Errorf("Error decoding: %s", err.Error())
  	}
  	daemonOSType := r.OSType
  
  	// Add labels to tags
  	for k, label := range container.Labels {
  		if d.labelFilter.Match(k) {
  			tags[k] = label
  		}
  	}
  
  	info, err := d.client.ContainerInspect(ctx, container.ID)
  	if err != nil {
  		return fmt.Errorf("Error inspecting docker container: %s", err.Error())
  	}
}
````
