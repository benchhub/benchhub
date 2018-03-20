package job

import (
	"fmt"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
)

// Planner generates execution plan based on assigned node and config
type Planner struct {
	log *dlog.Logger
}

func NewPlanner() *Planner {
	p := &Planner{}
	dlog.NewStructLogger(log, p)
	return p
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
	// TODO: we should support implicit pipeline in stage as we did for task, no pipeline means simply following defined order and execute one by one
	// expand stages in pipeline
	for _, pSpec := range jobSpec.Pipelines {
		pipeline := pb.StagePipelinePlan{
			Name: pSpec.Name,
		}
		// get all stages in this pipeline
		for _, stageName := range pSpec.Stages {
			stageSpec, ok := stageSpecs[stageName]
			if !ok {
				merr.Append(errors.Errorf("stage %s in pipeline %s not found", stageName, pSpec.Name))
				continue
			}
			// expand stage
			stage, err := p.Stage(nodes, stageSpec)
			if err != nil {
				merr.Append(err)
				continue
			}
			pipeline.Stages = append(pipeline.Stages, stage)
		}
		job.Pipelines = append(job.Pipelines, pipeline)
	}
	return job, merr.ErrorOrNil()
}

func (p *Planner) Stage(nodes []pb.AssignedNode, stageSpec pb.StageSpec) (pb.StagePlan, error) {
	merr := errors.NewMultiErr()
	// select nodes based on selectors, it's OR between selectors
	selection := make([]bool, len(nodes))
	selectedNodes := make([]pb.AssignedNode, 0, len(stageSpec.Selectors))
	for _, selector := range stageSpec.Selectors {
		for i, node := range nodes {
			if selection[i] {
				continue
			}
			if selector.Name == node.Spec.Name || selector.Role == node.Spec.Role {
				selectedNodes = append(selectedNodes, node)
				selection[i] = true
			}
		}
	}
	p.log.Infof("selected %d nodes from total %d nodes based on %d specs",
		len(selectedNodes), len(nodes), len(stageSpec.Selectors))

	// expand tasks in pipeline,
	var pipelines []pb.TaskPipelinePlan
	if len(stageSpec.Pipelines) == 0 {
		log.Infof("stage %s does not have explicit pipeline, use definition order", stageSpec.Name)
		for i, taskSpec := range stageSpec.Tasks {
			task, err := p.Task(nodes, taskSpec)
			if err != nil {
				merr.Append(err)
				continue
			}
			pipelines = append(pipelines, pb.TaskPipelinePlan{
				Name:  fmt.Sprintf("autogen-%d", i),
				Tasks: []pb.TaskPlan{task},
			})
		}
	} else {
		// pipeline is specified
		// TODO: check dup in pipeline spec
		taskSpecs, err := tasksToMap(stageSpec.Tasks)
		if err != nil {
			merr.Append(err)
			// NOTE: we don't stop here because the error can only be duplicated name in tasks
		}
		for _, pSpec := range stageSpec.Pipelines {
			// get all tasks in this pipeline
			pipeline := pb.TaskPipelinePlan{Name: pSpec.Name}
			for _, taskName := range pSpec.Tasks {
				taskSpec, ok := taskSpecs[taskName]
				if !ok {
					merr.Append(errors.Errorf("task %s in pipeline %s not found", taskName, pSpec.Name))
					continue
				}
				// expand task
				task, err := p.Task(nodes, taskSpec)
				if err != nil {
					merr.Append(err)
					continue
				}
				pipeline.Tasks = append(pipeline.Tasks, task)
			}
		}
	}

	// check if background flag is valid
	// a background stage must have at least one background task, and a non background stage should not
	// have any background task
	hasBackground := false
	for _, p := range pipelines {
		for _, t := range p.Tasks {
			if t.Spec.Background && !stageSpec.Background {
				// TODO: not all task has name, might just auto generate them?
				merr.Append(errors.Errorf("stage %s is not background but contains background task %s",
					stageSpec.Name, t.Spec.Name))
			}
			if t.Spec.Background {
				hasBackground = true
			}
		}
	}
	if stageSpec.Background && !hasBackground {
		merr.Append(errors.Errorf("stage %s is marked as background but no background task is defined", ))
	}
	stage := pb.StagePlan{
		Nodes:     selectedNodes,
		Pipelines: pipelines,
	}
	return stage, merr.ErrorOrNil()
}

func (p *Planner) Task(nodes []pb.AssignedNode, taskSpec pb.TaskSpec) (pb.TaskPlan, error) {
	return pb.TaskPlan{Spec: taskSpec}, nil
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

func tasksToMap(tasks []pb.TaskSpec) (map[string]pb.TaskSpec, error) {
	merr := errors.NewMultiErr()
	m := make(map[string]pb.TaskSpec, len(tasks))
	for i, task := range tasks {
		// no tasks should have same name, unique in one stage
		if _, ok := m[task.Name]; ok {
			merr.Append(errors.Errorf("task %s is defined again in %d", task.Name, i))
		}
		m[task.Name] = task
	}
	return m, merr.ErrorOrNil()
}
