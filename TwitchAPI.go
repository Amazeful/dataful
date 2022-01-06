package dataful

import "github.com/Amazeful/helix"

type TwitchAPI interface {
	NewAPI(options *helix.Options) (*helix.Client, error)
}

type Helix struct {
	clientID     string
	clientSecret string
}

//NewHelix returns a new TwitchAPI client.
func NewHelix(clientID, clientSecret string) *Helix {
	return &Helix{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

//NewAPI provides a new twitch helix API client with given options.
//ClientID and secret will be added unless they are included in options.
func (h *Helix) NewAPI(options *helix.Options) (*helix.Client, error) {
	if options.ClientID == "" {
		options.ClientID = h.clientID
	}
	if options.ClientSecret == "" {
		options.ClientSecret = h.clientSecret
	}
	return helix.NewClient(options)
}
