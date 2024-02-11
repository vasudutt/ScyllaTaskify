package scylla

import (
	"github.com/gocql/gocql"
	"github.com/vasudutt/ScyllaTaskify/config"
)

type Manager struct {
	cfg config.Config
}

func NewManager(cfg config.Config) *Manager {
	return &Manager{cfg: cfg}
}

func (m *Manager) Connect() (*gocql.Session, error) {
	return connect(m.cfg.ScyllaKeyspace, m.cfg.ScyllaHosts)
}

func (m *Manager) CreateKeyspace(keyspace string) error {
	session, err := connect("system", m.cfg.ScyllaHosts)
	if err != nil {
		return err
	}

	defer session.Close()

	statement := "CREATE KEYSPACE IF NOT EXISTS " + keyspace + " WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}"

	return session.Query(statement).Exec()
}

func connect(keyspace string, hosts []string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace

	return cluster.CreateSession()
}
