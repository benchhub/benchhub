package spec

type Task struct {
	// Daemon decides if the task is a long running process
	Daemon bool `yaml:"daemon"`
}
