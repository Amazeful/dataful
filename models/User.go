package models

import (
	"context"

	"github.com/Amazeful/dataful"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	dataful.BaseModel `bson:",inline"`

	UserID          string             `bson:"userId" json:"userId"`
	Login           string             `bson:"login" json:"login"`
	DisplayName     string             `bson:"displayName" json:"displayName"`
	AccessToken     string             `bson:"accessToken" json:"-"`
	RefreshToken    string             `bson:"refreshToken" json:"-"`
	Type            string             `bson:"type" json:"type"`
	BroadcasterType string             `bson:"broadcasterType" json:"broadcasterType"`
	Description     string             `bson:"description" json:"description"`
	ProfileImageURL string             `bson:"profileImageURL" json:"profileImageURL"`
	OfflineImageURL string             `bson:"offlineImageURL" json:"offlineImageURL"`
	ViewCount       int                `bson:"viewCount" json:"viewCount"`
	Suspended       bool               `bson:"suspended" json:"suspended"`
	Admin           bool               `bson:"admin" json:"admin"`
	Channel         primitive.ObjectID `bson:"channel" json:"channel"`
}

func NewUser(r dataful.Repository) *User {
	return &User{
		BaseModel: dataful.NewBaseModel(r),
	}
}

func (u *User) LoadBylId(ctx context.Context, id primitive.ObjectID) error {
	return u.R().FindOne(ctx, bson.M{"_id": id}, u)
}

func (u *User) LoadByUserId(ctx context.Context, userId string) error {
	return u.R().FindOne(ctx, bson.M{"userId": userId}, u)

}

func (u *User) Create(ctx context.Context) error {
	return u.R().InsertOne(ctx, u)
}

func (u *User) Update(ctx context.Context) error {
	return u.R().ReplaceOne(ctx, bson.M{"_id": u.ID}, u)
}
