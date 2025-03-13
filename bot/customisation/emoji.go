package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1349776909008896010, false)
	EmojiOpen       = NewCustomEmoji("open", 1349776942735556669, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1349776926331375717, false)
	EmojiClose      = NewCustomEmoji("close", 1349776883859849216, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1349776784287203450, false)
	EmojiReason     = NewCustomEmoji("reason", 1349776822488928358, false)
	EmojiSubject    = NewCustomEmoji("subject", 1349776839173996594, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1349776867195883610, false)
	EmojiClaim      = NewCustomEmoji("claim", 1349776670172905563, false)
	EmojiPanel      = NewCustomEmoji("panel", 1349776687071756379, false)
	EmojiRating     = NewCustomEmoji("rating", 1349776703760760895, false)
	EmojiStaff      = NewCustomEmoji("staff", 1349776744164495420, false)
	EmojiThread     = NewCustomEmoji("thread", 1349776589457461351, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1349776510302818416, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1349776552136544387, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1349776388965662762, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
