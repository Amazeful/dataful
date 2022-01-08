package models

import (
	"context"

	"github.com/Amazeful/dataful"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommandList struct {
	List []*Command `json:"list"`

	r dataful.Repository
}

func NewCommandList(r dataful.Repository) *CommandList {
	return &CommandList{
		List: make([]*Command, 0),
		r:    r,
	}
}

//LoadAllByChannel gets all commands for given channel.
func (cl *CommandList) LoadAllByChannel(ctx context.Context, channel primitive.ObjectID) error {
	err := cl.r.FindAll(ctx, bson.M{"channel": channel}, cl.List)
	if err != nil {
		return err
	}
	cl.setLoaded()
	return nil
}

func (cl *CommandList) setLoaded() {
	for _, commond := range cl.List {
		commond.SetLoaded(true)
		commond.SetR(cl.r)
	}
}
