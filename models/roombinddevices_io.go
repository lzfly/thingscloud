package models

type RoomBindDevicePostForm struct {
    Room_name           string `json:"room_name"`
	Device_sn            string `json:"device_sn"`
	Device_name          string `json:"device_name"`
    Type_code            int32  `json:"type_code"`
	Gateway_sn           string `json:"gateway_sn"`
	Username             string `json:"username"`
}

type RoomBindDevicePostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomBindDeviceInfo *RoomBindDevice `json:"roombinddevice"`
}

type RoomBindDeviceGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomBindDeviceInfo *RoomBindDevice `json:"roombinddevice"`
}

type RoomBindDeviceGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomBindDevicesInfo []RoomBindDevice `json:"roombinddevices"`
}

type RoomBindDevicePutForm struct {
    Room_name           string `json:"room_name"`
	Device_sn            string `json:"device_sn"`
	Device_name          string `json:"device_name"`
    Type_code            int32  `json:"type_code"`
	Gateway_sn           string `json:"gateway_sn"`
	Username             string `json:"username"`
}

type RoomBindDevicePutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	RoomBindDeviceInfo *RoomBindDevice `json:"roombinddevice"`
}

type RoomBindDeviceDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
