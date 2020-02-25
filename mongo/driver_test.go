package mongo_test

import (
	"context"
	"testing"

	gm "github.com/marrbor/go-mongo/mongo"
	"github.com/marrbor/golog"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewDriver(t *testing.T) {
	d, err := gm.NewDriver("abcdefg")
	assert.EqualError(t, err, `error parsing uri: scheme must be "mongodb" or "mongodb+srv"`)
	assert.Nil(t, d)

	url := "mongodb://root:example@localhost:27017"

	d, err = gm.NewDriver(url)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	err = d.Start(context.Background())
	assert.NoError(t, err)

	col := d.Collection("foo", "bar")
	assert.NotNil(t, col)

	d.Stop(context.Background())
}

func TestSave(t *testing.T) {
	type testSave struct {
		ID   int    `bson:"id"`
		Name string `bson:"name"`
	}

	url := "mongodb://root:example@localhost:27017"
	db := "go-mongo-test"
	collection := "test"

	golog.Info("let's save")
	ts := testSave{ID: 12345, Name: "foo"}
	if err := gm.Save(db, collection, url, context.Background(), ts); err != nil {
		assert.NoError(t, err)
	}

	d, err := gm.NewDriver(url)
	assert.NoError(t, err)
	assert.NotNil(t, d)

	ctx := context.Background()
	err = d.Start(ctx)
	assert.NoError(t, err)

	col := d.Collection(db, collection)
	assert.NotNil(t, col)

	golog.Info("let's load")
	sr := col.FindOne(ctx, bson.D{}, options.FindOne())
	assert.NoError(t, sr.Err())

	// check result
	var ts2 testSave
	assert.NoError(t, sr.Decode(&ts2))
	assert.EqualValues(t, ts.ID, ts2.ID)
	assert.EqualValues(t, ts.Name, ts2.Name)

	d.Stop(context.Background())
}
