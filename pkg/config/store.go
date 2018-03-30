package config

type MetaStoreConfig struct {
	Provider string `yaml:"provider"`
}

type TimeSeriesStoreConfig struct {
	Provider string `yaml:"provider"`
}
