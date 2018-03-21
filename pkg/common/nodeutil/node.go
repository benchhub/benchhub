package nodeutil

import (
	"os"

	"github.com/dyweb/gommon/errors"

	"github.com/benchhub/benchhub/lib/monitor/host"
	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/common/config"
	"strings"
)

// return node info that is needed when register agent and heartbeat

// GetNode returns node id, capacity, start & boot time
// TODO: addr https://github.com/benchhub/benchhub/issues/18
func GetNodeInfo(cfg config.NodeConfig) (*pb.NodeInfo, error) {
	m := host.NewMachine()
	if err := m.Update(); err != nil {
		return nil, errors.Wrap(err, "can't get node info")
	}
	node := &pb.NodeInfo{
		Id:        UID(),
		Role:      NodeRole(cfg.Role),
		Host:      hostname(),
		BootTime:  int64(m.BootTime), // unix ts in second
		StartTime: startTime.Unix(),  // unix ts in second
		Capacity: pb.NodeCapacity{
			Cores:       int32(m.NumCores),
			MemoryFree:  int32(m.MemFree / 1024),              // KB -> MB
			MemoryTotal: int32(m.MemTotal / 1024),             // KB -> MB
			DiskFree:    int32(m.DiskSpaceFree / 1024 / 1024), // B -> KB -> MB
			DiskTotal:   int32(m.DiskSpaceTotal / 1024 / 1024),
		},
	}
	return node, nil
}

// FIXME: might use pb package in config directly to avoid this issue
func NodeRole(role string) pb.Role {
	s := strings.ToLower(role)
	switch s {
	case "any":
		return pb.Role_ANY
	case "central":
		return pb.Role_CENTRAL
	case "loader":
		return pb.Role_LOADER
	case "database":
		return pb.Role_DATABASE
	}
	return pb.Role_UNKNOWN_ROLE
}

func hostname() string {
	if name, err := os.Hostname(); err != nil {
		log.Warnf("can't get hostname %v", err)
		return "unknown"
	} else {
		return name
	}
}
