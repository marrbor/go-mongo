package mongo_test

import (
	"context"
	"testing"

	"github.com/marrbor/go-mongo/mongo"
	"github.com/stretchr/testify/assert"
)

func TestNewDriver(t *testing.T) {
	d, err := mongo.NewDriver("abcdefg")
	assert.EqualError(t, err, `error parsing uri: scheme must be "mongodb" or "mongodb+srv"`)
	assert.Nil(t, d)

	d, err = mongo.NewDriver("mongodb://localhost:27070")
	assert.NoError(t, err)
	assert.NotNil(t, d)

	err = d.Start(context.Background())
	assert.NoError(t, err)

	col := d.Collection("foo", "bar")
	assert.NotNil(t, col)

	d.Stop(context.Background())
}
