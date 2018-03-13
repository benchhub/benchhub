package job

import (
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// Planner generates execution plan based on assigned node and config
type Planner struct {
	log *dlog.Logger
}

// TODO: generate the execution plan
// TODO: proto for plan
func (p *Planner) Job(nodes []pb.AssignedNode, jobSpec pb.JobSpec) (pb.JobPlan, error) {
	merr := errors.NewMultiErr()
	job := pb.JobPlan{}
	// convert stage array to map
	stageSpecs, err := stagesToMap(jobSpec.Stages)
	if err != nil {
		return job, err
	}
	// expand stages in pipeline
	for _, pSpec := range jobSpec.Pipelines {
		pipeline := pb.StagePipelinePlan{
			Name: pSpec.Name,
		}
		// get all the stages in this pipeline
		for _, stageName := range pSpec.Stages {
			stageSpec, ok := stageSpecs[stageName]
			if !ok {
				merr.Append(errors.Errorf("stage %s in pipeline %s not found", stageName, pSpec.Name))
			}
			// expand stage
			stage, err := p.Stage(nodes, stageSpec)
			if err != nil {
				merr.Append(err)
			}
			pipeline.Stages = append(pipeline.Stages, stage)
		}
	}
	return job, merr.ErrorOrNil()
}

func (p *Planner) Stage(nodes []pb.AssignedNode, stageSpec pb.StageSpec) (pb.StagePlan, error) {
	merr := errors.NewMultiErr()
	stage := pb.StagePlan{}
	// TODO: expand task
	return stage, merr.ErrorOrNil()
}

func stagesToMap(stages []pb.StageSpec) (map[string]pb.StageSpec, error) {
	merr := errors.NewMultiErr()
	m := make(map[string]pb.StageSpec, len(stages))
	for i, stage := range stages {
		// no stage should have same name, unique in one job spec
		if _, ok := m[stage.Name]; ok {
			merr.Append(errors.Errorf("stage %s is defined again in %d", stage.Name, i))
		}
		m[stage.Name] = stage
	}
	return m, merr.ErrorOrNil()
}
