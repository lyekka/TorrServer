package state

import (
	"server/torr/state"
)

type CacheState struct {
	Hash         string
	Capacity     int64
	Filled       int64
	PiecesLength int64
	PiecesCount  int
	Torrent      *state.TorrentStatus
	Pieces       map[int]ItemState
}

type ItemState struct {
	Id         int
	Length     int64
	Size       int64
	Completed  bool
	ReaderType int
}
