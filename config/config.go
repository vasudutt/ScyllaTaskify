package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ScyllaHosts         []string `envconfig:"SCYLLA_HOSTS" default:"127.0.0.1"`
	ScyllaKeyspace      string   `envconfig:"SCYLLA_KEYSPACE" default:"mykeyspace"`
	ScyllaMigrationsDir string   `envconfig:"SCYLLA_MIGRATIONS_DIR" default:"migrations"`
}

func Load() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return cfg, err
}
