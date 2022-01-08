package dataful

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Model interface.
type Model interface {
	Created()
	Updated()
	SetID(id interface{})
	SetLoaded(loaded bool)
}

//BaseModel includes fields shared by all models.
type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	r        Repository
	isLoaded bool
}

//NewBaseModel provides a new base model with given repository.
func NewBaseModel(r Repository) BaseModel {
	return BaseModel{r: r}
}

//R returns model repository.
func (bm *BaseModel) R() Repository {
	return bm.r
}

//Created populates createdat and updatedat fields.
func (bm *BaseModel) Created() {
	bm.CreatedAt = time.Now().UTC()
	bm.UpdatedAt = time.Now().UTC()
}

//Updated populates updatedat fields.
func (bm *BaseModel) Updated() {
	bm.UpdatedAt = time.Now().UTC()
}

//SetId sets the object id field of the model.
func (bm *BaseModel) SetID(id interface{}) {
	bm.ID = id.(primitive.ObjectID)
}

//SetLoaded sets a model's loaded flag.
func (bm *BaseModel) SetLoaded(loaded bool) {
	bm.isLoaded = loaded
}

//SetR sets a model's repository.
//Only use this in model list.
func (bm *BaseModel) SetR(r Repository) {
	bm.r = r
}

//Loaded returns a flag indicating if the model was successfully loaded from db.
func (bm *BaseModel) Loaded() bool {
	return bm.isLoaded
}
