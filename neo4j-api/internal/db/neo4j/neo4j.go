package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4j struct {
	Driver neo4j.DriverWithContext
}

func NewNeo4j(uri, username, password string) (*Neo4j, error) {
	driver, err := connect(uri, username, password)
	if err != nil {
		return nil, err
	}

	return &Neo4j{
		Driver: driver,
	}, nil
}

func connect(uri, username, password string) (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
}

func (r *Neo4j) Close(ctx context.Context) error {
	return r.Driver.Close(ctx)
}
