package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kashifkhan0771/TodoApp/config"
	"github.com/kashifkhan0771/TodoApp/db"
	domainErr "github.com/kashifkhan0771/TodoApp/errors"
	"github.com/kashifkhan0771/TodoApp/models"
)

const (
	taskCollection = "tasks"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mongo database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s/?connect=direct", viper.GetString(config.DBHost), viper.GetString(config.DBPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) AddTask(ctx context.Context, task *models.Task) (string, error) {
	if task.ID == "" {
		task.ID = uuid.NewV4().String()
	}

	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(taskCollection)
	if _, err := collection.InsertOne(ctx, task); err != nil {
		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

func (c *client) DeleteTask(ctx context.Context, id string) error {
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(taskCollection)
	if _, err := collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

func (c *client) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	var task *models.Task
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(taskCollection)
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}
	}

	return task, nil
}

func (c *client) UpdateTask(ctx context.Context, task *models.Task) error {
	collection := c.conn.Database(viper.GetString(config.DBName)).Collection(taskCollection)
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": task.ID}, bson.M{"$set": task}); err != nil {
		return errors.Wrap(err, "failed to update host")
	}

	return nil
}
