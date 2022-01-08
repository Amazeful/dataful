package models

import (
	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models/embeddables"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Purge struct {
	dataful.BaseModel `bson:",inline"`
	Channel           primitive.ObjectID `bson:"channel" json:"channel"`

	Enabled        bool                 `bson:"enabled" json:"enabled"`
	MaxRole        embeddables.UserRole `bson:"maxRole" json:"maxRole"`
	Lookback       int                  `bson:"lookback" json:"lookback"`
	Continuous     bool                 `bson:"continuous" json:"continuous"`
	ContinuousTime int                  `bson:"continuousTime" json:"continuousTime"`
}

func NewPurge(r dataful.Repository) *Purge {
	return &Purge{
		BaseModel:      dataful.NewBaseModel(r),
		MaxRole:        embeddables.UserRoleModerator,
		Lookback:       180,
		Continuous:     false,
		ContinuousTime: 60,
	}
}
