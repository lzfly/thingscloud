package models

type UserBindDevicePostForm struct {
	Username           string `json:"username"`
	Device_sn          string `json:"device_sn"`
	Gateway_sn         string `json:"gateway_sn"`
}

type UserBindDevicePostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindDeviceInfo *UserBindDevice `json:"userbinddevice"`
}

type UserBindDeviceGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindDeviceInfo *UserBindDevice `json:"userbinddevice"`
}

type UserBindDeviceGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindDevicesInfo []UserBindDevice `json:"userbinddevices"`
}

type UserBindDevicePutForm struct {
	Username           string  `json:"username"`
	Device_sn          string `json:"device_sn"`
	Gateway_sn         string `json:"gateway_sn"`
}

type UserBindDevicePutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindDeviceInfo *UserBindDevice `json:"userbinddevice"`
}

type UserBindDeviceDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
