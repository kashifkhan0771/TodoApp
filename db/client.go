package db

import (
	"context"
	"log"

	"github.com/kashifkhan0771/TodoApp/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	AddTask(ctx context.Context, t *models.Task) (string, error)
	DeleteTask(ctx context.Context, id string) error
	GetTaskByID(ctx context.Context, id string) (*models.Task, error)
	UpdateTask(ctx context.Context, t *models.Task) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)

		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)

		return
	}
	datastoreFactories[name] = factory
}
