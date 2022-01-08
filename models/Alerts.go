package models

import (
	"github.com/Amazeful/dataful"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alerts struct {
	dataful.BaseModel `bson:",inline"`
	Channel           primitive.ObjectID `bson:"channel" json:"channel"`

	Sub           Alert `bson:"sub" json:"sub"`
	Resub         Alert `bson:"resub" json:"resub"`
	SubGift       Alert `bson:"subGift" json:"subGift"`
	CommunityGift Alert `bson:"communityGift" json:"communityGift"`
	Follow        Alert `bson:"follow" json:"follow"`
	Raid          Alert `bson:"raid" json:"raid"`
	Bits          Alert `bson:"bits" json:"bits"`
}

func NewAlerts(r dataful.Repository) *Alerts {
	return &Alerts{
		BaseModel:     dataful.NewBaseModel(r),
		Sub:           Alert{Message: "@$(user), Thank you for $(tier) subscription."},
		Resub:         Alert{Message: "@$(user), Thank you for subscribing for $(months) months in a row."},
		SubGift:       Alert{Message: "@$(gifter), Thank you for gifting a $(tier) subscription to $(user)."},
		CommunityGift: Alert{Message: "@$(gifter), Thank you for gifting $(count) tier $(tier) subs to $(channel)'s community."},
		Follow:        Alert{Message: "@$(user), Thank you for following."},
		Raid:          Alert{Message: "@$(user), Thank you for raiding the stream with $(viewers) viewers."},
		Bits:          Alert{Message: "@$(user), Thank you for $(bits) bits."},
	}
}

type Alert struct {
	Enabled bool   `bson:"enabled" json:"enabled"`
	Message string `bson:"message" json:"message"`
}
