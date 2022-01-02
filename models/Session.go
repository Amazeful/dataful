package models

import (
	"context"
	"time"

	"github.com/Amazeful/dataful"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	SessionId       string             `json:"sessionId"`
	User            primitive.ObjectID `json:"user"`
	SelectedChannel primitive.ObjectID `json:"channel"`

	c dataful.Cache
}

func NewSession(c dataful.Cache) *Session {
	return &Session{c: c}
}

//GenerateSessionId generates a new session uid
func (s *Session) GenerateSessionId() {
	s.SessionId = xid.New().String()
}

//SetSession adds session to db
func (s *Session) SetSession(ctx context.Context, expiry time.Duration) error {
	return s.c.Set(ctx, s.key(), s, expiry)
}

//GetSession gets the session from db
func (s *Session) GetSession(ctx context.Context) error {
	return s.c.Get(ctx, s.key(), s)
}

func (s *Session) key() string {
	return "session-" + s.SessionId
}
