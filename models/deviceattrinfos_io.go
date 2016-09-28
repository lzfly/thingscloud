package models

type DeviceAttrInfoPostForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_name            string  `json:"attr_name"`
	Attr_permission      string  `json:"attr_permission"`
	Attr_value_ctrl      string  `json:"attr_value_ctrl"`
	Is_control           int32  `json:"is_control"`
	Attr_value_cur       string  `json:"attr_value_cur"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceAttrInfoPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfoInfo *DeviceAttrInfo `json:"deviceattrinfo"`
}

type DeviceAttrInfoGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfoInfo *DeviceAttrInfo `json:"deviceattrinfo"`
}

type DeviceAttrInfoGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfosInfo []DeviceAttrInfo `json:"deviceattrinfos"`
}

type DeviceAttrInfoPutForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_name            string  `json:"attr_name"`
	Attr_permission      string  `json:"attr_permission"`
	Attr_value_ctrl      string  `json:"attr_value_ctrl"`
	Is_control           int32  `json:"is_control"`
	Attr_value_cur       string  `json:"attr_value_cur"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceAttrInfoPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfoInfo *DeviceAttrInfo `json:"deviceattrinfo"`
}

type DeviceAttrInfoDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
