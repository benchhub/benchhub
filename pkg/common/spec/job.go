package spec

type Job struct {
	Loader   Loader   `yaml:"loader"`
	Database Database `yaml:"database"`
}
