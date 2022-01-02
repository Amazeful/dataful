package models

import (
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	type args struct {
		signKey    string
		expiration time.Time
		sessionId  string
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{"valid", false, args{"test", time.Now(), "123"}},
		{"valid2", false, args{"test", time.Now(), "321"}},
		{"invalid sign key", true, args{"", time.Now(), "321"}},
		{"invalid time", true, args{"", time.Now().Add(-time.Hour), "321"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			j := NewJWT([]byte(test.args.signKey), jwa.HS256)
			tokenString, err := j.Encode(test.args.sessionId, test.args.expiration)
			if test.wantErr {
				assert.Error(t, err)
				assert.Empty(t, tokenString)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, tokenString)

				token, err := jwt.Parse([]byte(tokenString))
				assert.NoError(t, err)
				assert.Equal(t, test.args.expiration.Unix(), token.Expiration().Unix())
				sid, ok := token.Get(SessionIdKey)
				assert.True(t, ok)
				assert.Equal(t, test.args.sessionId, sid)
			}

		})
	}

}

func TestDecode(t *testing.T) {

	type args struct {
		signKey    string
		decodeKey  string
		expiration time.Time
		sessionId  string
		algorithm  jwa.SignatureAlgorithm
	}

	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{"valid", false, args{"test", "test", time.Now().Add(time.Minute), "123", jwa.HS256}},
		{"valid2", false, args{"test", "test", time.Now().Add(time.Minute), "321", jwa.HS256}},
		{"invalid sign key", true, args{"test", "tset", time.Now().Add(time.Minute), "321", jwa.HS256}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token, err := jwt.NewBuilder().
				IssuedAt(time.Now().UTC()).
				Expiration(test.args.expiration).
				Claim(SessionIdKey, test.args.sessionId).
				Build()
			assert.NoError(t, err)

			payload, err := jwt.Sign(token, test.args.algorithm, []byte(test.args.signKey))
			assert.NoError(t, err)

			j := NewJWT([]byte(test.args.decodeKey), test.args.algorithm)
			result, err := j.Decode(string(payload))
			if test.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				sid, ok := result.Get(SessionIdKey)
				assert.True(t, ok)
				assert.Equal(t, test.args.sessionId, sid)
			}

		})
	}

}
