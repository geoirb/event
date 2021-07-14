package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/geoirb/event/pkg/model"
	"github.com/geoirb/event/pkg/mongo"
)

// Builder query for mongo storage.
type Builder struct{}

// NewBuilder ...
func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) NewEvent(e model.Event) mongo.Event {
	return mongo.Event{
		Type:       e.Type,
		State:      int64(e.State),
		StartedAt:  e.StartedAt,
		FinishedAt: e.FinishedAt,
	}
}

func (b *Builder) UpdateEvent(e model.Event) (filter bson.M, update bson.M) {
	filter, update = make(primitive.M), make(primitive.M)

	filter["type"] = e.Type

	update["finished_at"] = e.FinishedAt
	update["state"] = e.State
	return
}
