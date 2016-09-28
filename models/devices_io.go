package models

type DevicePostForm struct {
	Device_sn      string `json:"device_sn"`
	Type_code      int32  `json:"type_code"`
	Type_name      string `json:"type_name"`
	Device_model   string `json:"device_model"`
	Device_ver     string `json:"device_ver"`
	Protocol       string `json:"protocol"`
	Device_name           string `json:"device_name"`
	Gateway_sn     string `json:"gateway_sn"`
}

type DevicePostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceInfo *Device `json:"device"`
}

type DeviceGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceInfo *Device `json:"device"`
}

type DeviceGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DevicesInfo []Device2 `json:"devices"`
}

type DevicePutForm struct {
	Type_code      int32  `json:"type_code"`
	Type_name      string `json:"type_name"`
	Device_model   string `json:"device_model"`
	Device_ver     string `json:"device_ver"`
	Protocol       string `json:"protocol"`
	Device_name           string `json:"device_name"`
	Gateway_sn     string `json:"gateway_sn"`
}

type DevicePutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceInfo *Device `json:"device"`
}

type DeviceDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
