package model

type Room struct {
	Rid             int
	Mid             int
	State           int
	Capacity        int
	PlayerCount     int
	RoomName        string
	OwnerName       string
	GeneralSetting  int
	MaxCampCount    int
	MaxWatcherCount int
	MapID           int
	WeatherType     int
	LightType       int
	IfPublic        int
	Password        string
	Type            string
}

type RoomListInfo struct {
	OwnerName   string
	RoomName    string
	Capacity    string
	PlayerCount string
	IfPublic    string
	State       string
}

type RoomConfig struct {
	Rid         int
	OwnerName   string
	Weather     string
	Time        string
	Map         string
	PlayerCount string
	State       int
}

type Begin struct {
	OperationType int32
	DaotiaoType   int32
	Timestamp     int32
	ContentLength int32
	CallbackID    int32
	Room          Room
	UsersInfo     []*RoomPlayer
}
