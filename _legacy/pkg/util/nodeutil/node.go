package nodeutil

import (
	"os"

	"github.com/dyweb/gommon/errors"

	"github.com/benchhub/benchhub/lib/monitor/host"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// GetNode returns node id, capacity, start & boot time that is needed when register agent and heartbeat
func GetNodeInfo(cfg pb.NodeConfig, bindAddr string) (*pb.NodeInfo, error) {
	m := host.NewMachine()
	if err := m.Update(); err != nil {
		return nil, errors.Wrap(err, "can't get node info")
	}
	node := &pb.NodeInfo{
		Id: UID(),
		Addr: pb.Addr{
			BindAddr: bindAddr,
			// TODO: get addr https://github.com/benchhub/benchhub/issues/18
		},
		Config: cfg,
		Capacity: pb.NodeCapacity{
			Cores:       int32(m.NumCores),
			MemoryFree:  int32(m.MemFree / 1024), // KB -> MB
			MemoryTotal: int32(m.MemTotal / 1024),
			DiskFree:    int32(m.DiskSpaceFree / 1024 / 1024), // B -> KB -> MB
			DiskTotal:   int32(m.DiskSpaceTotal / 1024 / 1024),
		},
		Property: pb.NodeProperty{
			BootTime:  int64(m.BootTime), // unix ts in second
			StartTime: startTime.Unix(),  // unix ts in second
			Host:      Hostname(),
		},
	}
	return node, nil
}

func Hostname() string {
	if name, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return name
	}
}
