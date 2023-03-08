package chunk

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

// ChunkType is the type of chunk that was recorded.  There are three types:
// messages, thread messages, and files.
type ChunkType int

//go:generate stringer -type=ChunkType -trimprefix=C
const (
	CMessages ChunkType = iota
	CThreadMessages
	CFiles
	CUsers
	CChannels
	CChannelInfo
)

// Chunk is a single chunk that was recorded.  It contains the type of chunk,
// the timestamp of the chunk, the channel ID, and the number of messages or
// files that were recorded.
type Chunk struct {
	Type      ChunkType `json:"_t"`
	Timestamp int64     `json:"_ts"`
	IsThread  bool      `json:"_tm,omitempty"`
	Count     int       `json:"_c"` // number of messages or files

	// Channel contains the channel information.  It may not be immediately
	// followed by messages from the channel.  Populated by ChannelInfo method.
	Channel *slack.Channel `json:"_ci,omitempty"`

	ChannelID string `json:"_id"`
	// Parent is populated in case the chunk is a thread, or a file. Populated
	// by ThreadMessages and Files methods.
	Parent *slack.Message `json:"_p,omitempty"`
	// Messages contains a chunk of messages as returned by the API. Populated
	// by Messages and ThreadMessages methods.
	Messages []slack.Message `json:"_m,omitempty"`
	// Files contains a chunk of files as returned by the API. Populated by
	// Files method.
	Files []slack.File `json:"_f,omitempty"`

	// Users contains a chunk of users as returned by the API. Populated by
	// Users method.
	Users []slack.User `json:"_u,omitempty"` // Populated by Users
	// Channels contains a chunk of channels as returned by the API. Populated
	// by Channels method.
	Channels []slack.Channel `json:"_ch,omitempty"` // Populated by Channels
}

const (
	userChunkID    = "usr"
	channelChunkID = "chan"
)

// ID returns a unique ID for the chunk.
func (c *Chunk) ID() string {
	switch c.Type {
	case CMessages:
		return c.ChannelID
	case CThreadMessages:
		return threadID(c.ChannelID, c.Parent.ThreadTimestamp)
	case CFiles:
		return id("f", c.ChannelID, c.Parent.Timestamp)
	case CChannelInfo:
		return channelInfoID(c.ChannelID, c.IsThread)
	case CUsers:
		return userChunkID // static, one team per chunk file
	case CChannels:
		return channelChunkID // static, one team per chunk file.
	}
	return fmt.Sprintf("<unknown:%s>", c.Type)
}

func id(prefix string, ids ...string) string {
	return prefix + strings.Join(ids, ":")
}

func threadID(channelID, threadTS string) string {
	return id("t", channelID, threadTS)
}

func channelInfoID(channelID string, isThread bool) string {
	if isThread {
		return id("tci", channelID)
	}
	return id("ci", channelID)
}
