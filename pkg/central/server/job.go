package server

import (
	"context"
	"time"

	dlog "github.com/dyweb/gommon/log"

	pbc "github.com/benchhub/benchhub/pkg/common/commonpb"

	"github.com/benchhub/benchhub/pkg/common/spec"
	"github.com/dyweb/gommon/errors"
)

type JobController struct {
	registry *Registry
	log      *dlog.Logger
}

type AssignResult struct {
	Spec spec.Node
	Node pbc.Node
}

func NewJobController(r *Registry) (*JobController, error) {
	j := &JobController{
		registry: r,
	}
	dlog.NewStructLogger(log, j)
	return j, nil
}

func (j *JobController) RunWithContext(ctx context.Context) error {
	j.log.Info("start job controller")
	meta := j.registry.Meta
	for {
		select {
		case <-ctx.Done():
			// TODO: should we return nil or return context error?
			// TODO: we should tell all the agent to abort job since central is shut down?
			j.log.Infof("job controller stop due to context finished, its error is %v", ctx.Err())
			return nil
		default:
			job, empty, err := meta.GetPendingJob()
			if empty {
				// do nothing
			} else if err != nil {
				log.Warnf("failed to get pending job %v", err)
			} else {
				// TODO: spec should contains id ...
				nodes, err := meta.ListNodes()
				if err != nil {
					log.Warnf("can't list nodes %s", err.Error())
				}
				// TODO: acquire node resource based on node selector
				results, err := j.AcquireNodes(nodes, job.Nodes)
				if err != nil {
					log.Warnf("can't acquire nodes %s", err.Error())
				}
				for _, r := range results {
					// TODO: print it or? ...
					log.Infof("result is %v", r)
				}
				log.Infof("TODO: process job %s", job.Name)
			}
			// TODO: poll duration should be configurable
			time.Sleep(1 * time.Second)
		}
	}
}

func (j *JobController) AcquireNodes(nodes []pbc.Node, specs []spec.Node) ([]AssignResult, error) {
	if len(nodes) == 0 {
		return nil, errors.New("0 agent no node to acquire")
	}
	if len(nodes) < len(specs) {
		j.log.Warnf("only %d agents but want %d nodes", len(nodes), len(specs))
	}
	used := make([]int, len(nodes))
	acquired := make([]int, len(specs))
	// first loop, don't reuse any node

	for i, s := range specs {
		for j, node := range nodes {
			if used[j] > 0 {
				continue
			}
			if (s.Type == spec.NodeTypeDatabase && node.Role == pbc.Role_DATABASE) ||
				(s.Type == spec.NodeTypeLoader && node.Role == pbc.Role_LOADER) {
				// NOTE: +1 so we can check if the spec has acquired node with > 0
				acquired[i] = j + 1
				used[j]++
				break
			}
			if node.Role == pbc.Role_ANY {
				acquired[i] = j + 1
				used[j]++
				break
			}
		}
	}
	// second loop,
	// TODO: allow multiple workload on node, but never assign database with loader on same node
	for maxUse := 0; maxUse < 3; maxUse++ {
		for i := range specs {
			if acquired[i] > 0 {
				continue
			}
			for j := range nodes {
				if used[j] > maxUse {
					continue
				}
				acquired[i] = j + 1
				used[j]++
				break
			}
		}
	}
	// all spec assigned to nodes?
	merr := errors.NewMultiErr()
	results := make([]AssignResult, 0, len(specs))
	for i, spc := range specs {
		if acquired[i] > 0 {
			j.log.Infof("plan: assign %s %s to node %s", spc.Name, spc.Type, nodes[acquired[i]-1].Uid)
			results = append(results, AssignResult{
				Spec: spc,
				Node: nodes[acquired[i]-1],
			})
		} else {
			merr.Append(errors.Errorf("spec %s %s has no node", spc.Name, spc.Type))
		}
	}
	return results, merr.ErrorOrNil()
}
