package dataful

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	Created()
	Updated()
	SetId(id interface{})
	SetLoaded(loaded bool)
}

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	r        Repository
	isLoaded bool
}

func NewBaseModel(r Repository) BaseModel {
	return BaseModel{r: r}
}

func (bm *BaseModel) R() Repository {
	return bm.r
}

func (bm *BaseModel) Created() {
	bm.CreatedAt = time.Now().UTC()
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) Updated() {
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) SetId(id interface{}) {
	bm.ID = id.(primitive.ObjectID)
}

func (bm *BaseModel) SetLoaded(loaded bool) {
	bm.isLoaded = loaded
}

func (bm *BaseModel) Loaded() bool {
	return bm.isLoaded
}
