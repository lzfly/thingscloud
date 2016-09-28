package models

type RoomPostForm struct {
	Room_name      string `json:"room_name"`
	Username        string `json:"username"`
}

type RoomPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomInfo *Room `json:"room"`
}

type RoomGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomInfo *Room `json:"room"`
}

type RoomGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomsInfo []Room `json:"rooms"`
}

type RoomPutForm struct {
	Room_name      string `json:"room_name"`
	Username        string `json:"username"`
}

type RoomPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomInfo *Room `json:"room"`
}

type RoomDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
