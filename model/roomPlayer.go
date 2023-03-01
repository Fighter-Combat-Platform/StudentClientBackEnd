package model

type RoomPlayer struct {
	Rid         int
	Uid         int
	GroupType   int
	FighterType int
	Ready       int
}

type RoomPlayerInfo struct {
	Uid         int
	PlayerName  string
	GroupType   string
	FighterType string
	Ready       string
}
