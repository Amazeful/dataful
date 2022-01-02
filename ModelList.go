package dataful

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelList struct {
	List []Model

	r Repository
}

func NewModelList(r Repository, models []Model) *ModelList {
	return &ModelList{
		r:    r,
		List: models,
	}
}

func (ml *ModelList) R() Repository {
	return ml.r
}

func (ml *ModelList) FindAll(ctx context.Context, filter bson.M, opts ...*options.FindOptions) error {
	return ml.R().FindAll(ctx, filter, ml.List, opts...)
}
