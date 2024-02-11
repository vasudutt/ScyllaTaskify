package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("SCYLLA_HOSTS", "host1.com,host2.com")
	os.Setenv("SCYLLA_MIGRATIONS_DIR", "./cql")
	os.Setenv("SCYLLA_KEYSPACE", "reporting")

	actual, err := Load()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := Config{
		ScyllaHosts:         []string{"host1.com", "host2.com"},
		ScyllaKeyspace:      "reporting",
		ScyllaMigrationsDir: "./cql",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	os.Clearenv()
}
