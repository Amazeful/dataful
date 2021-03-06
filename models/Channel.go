package models

import (
	"context"
	"time"

	"github.com/Amazeful/dataful"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Channel struct {
	dataful.BaseModel `bson:",inline"`

	ChannelId       string    `bson:"channelId" json:"channelId"`
	BroadcasterName string    `bson:"broadcasterName" json:"broadcasterName"`
	Language        string    `bson:"language" json:"language"`
	GameId          string    `bson:"gameId" json:"gameId"`
	GameName        string    `bson:"gameName" json:"gameName"`
	Title           string    `bson:"title" json:"title"`
	Joined          bool      `bson:"joined" json:"joined"`
	Silenced        bool      `bson:"silenced" json:"silenced"`
	AccessToken     string    `bson:"accessToken,omitempty" json:"-"`
	RefreshToken    string    `bson:"refreshToken,omitempty" json:"-"`
	Prefix          string    `bson:"prefix" json:"prefix"`
	Live            bool      `bson:"live" json:"live"`
	Shard           int       `bson:"shard" json:"shard"`
	Authenticated   bool      `bson:"authenticated" json:"authenticated"`
	StartedAt       time.Time `bson:"startedAt,omitempty" json:"startedAt,omitempty"`
	EndedAt         time.Time `bson:"endedAt,omitempty" json:"endedAt,omitempty"`
	Moderator       bool      `bson:"moderator" json:"moderator"`
}

func NewChannel(r dataful.Repository) *Channel {
	return &Channel{
		BaseModel: dataful.NewBaseModel(r),
		Joined:    true,
		Prefix:    "!",
	}
}

func (c *Channel) LoadBylId(ctx context.Context, id primitive.ObjectID) error {
	return c.R().FindOne(ctx, bson.M{"_id": id}, c)
}

func (c *Channel) LoadByChannelName(ctx context.Context, name string) error {
	return c.R().FindOne(ctx, bson.M{"broadcasterName": name}, c)
}

func (c *Channel) LoadByChannelId(ctx context.Context, channelId string) error {
	return c.R().FindOne(ctx, bson.M{"channelId": channelId}, c)
}

func (c *Channel) Create(ctx context.Context) error {
	return c.R().InsertOne(ctx, c)
}

func (c *Channel) Update(ctx context.Context) error {
	return c.R().ReplaceOne(ctx, bson.M{"_id": c.ID}, c)
}

func (c *Channel) Delete(ctx context.Context) error {
	return c.R().DeleteOne(ctx, bson.M{"_id": c.ID})
}
