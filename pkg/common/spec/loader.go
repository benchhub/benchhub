package spec

type Loader struct {
	// Setup install/check required software to run the loader, i.e. jvm
	Setup []Task `yaml:"setup"`
	// Ping make sures all loader can reach database(s)
	Ping []Task `yaml:"ping"`
	// Migration creates database/table needed for loading
	Migration []Task `yaml:"migration"`
	// TODO: for daemon ... how to specify kill a long running task previously defined
	Shutdown []Task `yaml:"shutdown"`
}
