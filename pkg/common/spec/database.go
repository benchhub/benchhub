package spec

type Database struct {
	Setup []Task `yaml:"setup"`
	// TODO: how to handle server like KairosDB, which requires another database ...?
	Run []Task `yaml:"run"`
}
