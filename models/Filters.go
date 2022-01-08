package models

import (
	"context"

	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models/embeddables"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Filters struct {
	dataful.BaseModel `bson:",inline"`
	Channel           primitive.ObjectID `bson:"channel" json:"channel"`

	Caps       CapsFilter       `bson:"caps" json:"caps"`
	Symbols    SymbolsFilter    `bson:"symbols" json:"symbols"`
	Zalgo      ZalgoFilter      `bson:"zalgo" json:"zalgo"`
	English    EnglishFilter    `bson:"english" json:"english"`
	Length     LengthFilter     `bson:"length" json:"length"`
	Emotes     EmotesFilter     `bson:"emotes" json:"emotes"`
	Emojis     EmojisFilter     `bson:"emojis" json:"emojis"`
	SoloSpam   SoloSpamFilter   `bson:"soloSpam" json:"soloSpam"`
	Repetition RepetitionFilter `bson:"repetition" json:"repetition"`
	Link       LinkFilter       `bson:"link" json:"link"`
}

func NewFilters(r dataful.Repository) *Filters {
	return &Filters{
		BaseModel:  dataful.NewBaseModel(r),
		Caps:       NewCapsFilter(),
		Symbols:    NewSymbolsFilter(),
		Zalgo:      NewZalgoFilter(),
		English:    NewEnglishFilter(),
		Length:     NewLengthFilter(),
		Emotes:     NewEmotesFilter(),
		Emojis:     NewEmojisFilter(),
		SoloSpam:   NewSoloSpamFilter(),
		Repetition: NewRepetitionFilter(),
		Link:       NewLinkFilter(),
	}
}

func (f *Filters) LoadByChannel(ctx context.Context, channel primitive.ObjectID) error {
	return f.R().FindOne(ctx, bson.M{"channel": channel}, f)
}

func (f *Filters) Create(ctx context.Context) error {
	return f.R().InsertOne(ctx, f)
}

func (f *Filters) Update(ctx context.Context) error {
	return f.R().ReplaceOne(ctx, bson.M{"_id": f.ID}, f)
}

type FilterCommon struct {
	Enabled          bool                     `bson:"enabled" json:"enabled"`
	Description      string                   `bson:"description" json:"description"`
	Stream           embeddables.StreamStatus `bson:"stream" json:"stream"`
	DisableOnRaid    bool                     `bson:"disableOnRaid" json:"disableOnRaid"`
	Announce         bool                     `bson:"announce" json:"announce"`
	AnnounceCooldown int                      `bson:"announceCooldown" json:"announceCooldown"`
	Message          string                   `bson:"message" json:"message"`
	TimeoutDuration  int                      `bson:"timeoutDuration" json:"timeoutDuration"`
	ExponentialBase  int                      `bson:"exponentialBase" json:"exponentialBase"`
	OffenseTimeout   int                      `bson:"offenseTimeout" json:"offenseTimeout"`
	MaxRole          embeddables.UserRole     `bson:"maxRole" json:"maxRole"`
}

func NewFilterCommon(description, message string) FilterCommon {
	return FilterCommon{
		Description:      description,
		Stream:           embeddables.StreamLive | embeddables.StreamOffline,
		Announce:         true,
		AnnounceCooldown: 60,
		Message:          message,
		TimeoutDuration:  30,
		ExponentialBase:  1,
		OffenseTimeout:   180,
		MaxRole:          embeddables.UserRoleModerator,
	}
}

type CapsFilter struct {
	FilterCommon  `bson:",inline"`
	MinChars      int     `bson:"minChars" json:"minChars"`
	MaxCount      int     `bson:"maxCount" json:"maxCount"`
	MaxPercentage float64 `bson:"maxPercentage" json:"maxPercentage"`
}

func NewCapsFilter() CapsFilter {
	return CapsFilter{
		FilterCommon:  NewFilterCommon("Timeout users for using excessive CAPS.", "Please stop spamming caps."),
		MinChars:      50,
		MaxCount:      100,
		MaxPercentage: 0.9,
	}
}

type SymbolsFilter struct {
	FilterCommon  `bson:",inline"`
	MinChars      int     `bson:"minChars" json:"minChars"`
	MaxCount      int     `bson:"maxCount" json:"maxCount"`
	MaxPercentage float64 `bson:"maxPercentage" json:"maxPercentage"`
}

func NewSymbolsFilter() SymbolsFilter {
	return SymbolsFilter{
		FilterCommon:  NewFilterCommon("Timeout users for using excessive symbols.", "Please stop spamming symbols."),
		MinChars:      50,
		MaxCount:      100,
		MaxPercentage: 0.9,
	}
}

type ZalgoFilter struct {
	FilterCommon  `bson:",inline"`
	MinChars      int     `bson:"minChars" json:"minChars"`
	MaxCount      int     `bson:"maxCount" json:"maxCount"`
	MaxPercentage float64 `bson:"maxPercentage" json:"maxPercentage"`
}

func NewZalgoFilter() ZalgoFilter {
	return ZalgoFilter{
		FilterCommon:  NewFilterCommon("Timeout users for posting zalgo characters in chat.", "Please stop using zalgo characters."),
		MinChars:      1,
		MaxCount:      1,
		MaxPercentage: 0,
	}
}

type EnglishFilter struct {
	FilterCommon  `bson:",inline"`
	MinChars      int     `bson:"minChars" json:"minChars"`
	MaxCount      int     `bson:"maxCount" json:"maxCount"`
	MaxPercentage float64 `bson:"maxPercentage" json:"maxPercentage"`
}

func NewEnglishFilter() EnglishFilter {
	return EnglishFilter{
		FilterCommon:  NewFilterCommon("Timeout users for typing non-English characters in chat.", "English only please."),
		MinChars:      10,
		MaxCount:      5,
		MaxPercentage: 0.9,
	}
}

type LengthFilter struct {
	FilterCommon `bson:",inline"`
	MaxCount     int `bson:"maxCount" json:"maxCount"`
}

func NewLengthFilter() LengthFilter {
	return LengthFilter{
		FilterCommon: NewFilterCommon("Timeout users for posting lengthy messages in chat.", "Please stop posting lengthy messages."),
		MaxCount:     350,
	}
}

type EmotesFilter struct {
	FilterCommon `bson:",inline"`
	MaxCount     int `bson:"maxCount" json:"maxCount"`
}

func NewEmotesFilter() EmotesFilter {
	return EmotesFilter{
		FilterCommon: NewFilterCommon("Timeout users for using excessive Twitch emotes.", "Please stop spamming emotes."),
		MaxCount:     20,
	}
}

type EmojisFilter struct {
	FilterCommon `bson:",inline"`
	MaxCount     int `bson:"maxCount" json:"maxCount"`
}

func NewEmojisFilter() EmojisFilter {
	return EmojisFilter{
		FilterCommon: NewFilterCommon("Timeout users for using excessive emojis.", "Please stop spamming emojis."),
		MaxCount:     20,
	}
}

type SoloSpamFilter struct {
	FilterCommon `bson:",inline"`
	MinChars     int     `bson:"minChars" json:"minChars"`
	Similarity   float64 `bson:"similarity" json:"similarity"`
	MaxCount     int     `bson:"maxCount" json:"maxCount"`
	Lookback     int     `bson:"lookback" json:"lookback"`
}

func NewSoloSpamFilter() SoloSpamFilter {
	return SoloSpamFilter{
		FilterCommon: NewFilterCommon("Timeout users for solo spamming in chat.", "Please stop from solo spamming."),
		MinChars:     50,
		Similarity:   0.9,
		MaxCount:     3,
		Lookback:     180,
	}
}

type RepetitionFilter struct {
	FilterCommon   `bson:",inline"`
	MinChars       int `bson:"minChars" json:"minChars"`
	MaxRepetitions int `bson:"maxRepetitions" json:"maxRepetitions"`
	MinUnique      int `bson:"minUnique" json:"minUnique"`
}

func NewRepetitionFilter() RepetitionFilter {
	return RepetitionFilter{
		FilterCommon:   NewFilterCommon("Timeout users for repetitive words in their messages.", "Please stop repeating yourself."),
		MinChars:       50,
		MinUnique:      15,
		MaxRepetitions: 7,
	}
}

type LinkFilter struct {
	FilterCommon `bson:",inline"`
	Whitelist    []string `bson:"whitelist" json:"whitelist"`
}

func NewLinkFilter() LinkFilter {
	return LinkFilter{
		FilterCommon: NewFilterCommon("Timeout users for posting links in chat.", "Please stop posting links."),
		Whitelist:    []string{"twitch.tv"},
	}
}
