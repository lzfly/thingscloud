package models

type DeviceCtrlPostForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           int64  `json:"attr_value"`
	Source               string `json:"source"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceCtrlPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceCtrlInfo *DeviceCtrl `json:"devicectrl"`
}

type DeviceCtrlGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceCtrlInfo *DeviceCtrl `json:"devicectrl"`
}

type DeviceCtrlGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceCtrlsInfo []DeviceCtrl `json:"devicectrls"`
}

type DeviceCtrlPutForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           int64  `json:"attr_value"`
	Source               string `json:"source"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceCtrlPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceCtrlInfo *DeviceCtrl `json:"devicectrl"`
}

type DeviceCtrlDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
