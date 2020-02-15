package mongo

import (
	"context"

	"github.com/marrbor/golog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Driver struct {
	url    string
	client *mongo.Client
}

// Start connects the driver to mongo server.
func (d *Driver) Start(ctx context.Context) error {
	return d.client.Connect(ctx)
}

// Stop disconnects the driver to mongo server.
// Not return error (only logging) since easy to use for defer function.
func (d *Driver) Stop(ctx context.Context) {
	if err := d.client.Disconnect(ctx); err != nil {
		golog.Error(err)
	}
}

// Collection returns mongo collection ineterface based on given db and collection.
func (d *Driver) Collection(db, collection string) *mongo.Collection {
	return d.client.Database(db).Collection(collection)
}

// NewDriver returns new mongodb client.
func NewDriver(url string) (*Driver, error) {
	cl, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	return &Driver{url: url, client: cl,}, nil
}
