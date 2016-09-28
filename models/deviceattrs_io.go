package models

type DeviceAttrPostForm struct {
	Attr_name             string `json:"attr_name"`
	Attr_code        int32 `json:"attr_code"`
	Datatype         string `json:"datatype"`
}

type DeviceAttrPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfo *DeviceAttr `json:"deviceattr"`
}

type DeviceAttrGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfo *DeviceAttr `json:"deviceattr"`
}

type DeviceAttrGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrsInfo []DeviceAttr `json:"deviceattrs"`
}

type DeviceAttrPutForm struct {
	Attr_name             string `json:"attr_name"`
	Attr_code        int32 `json:"attr_code"`
	Datatype         string `json:"datatype"`
}

type DeviceAttrPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceAttrInfo *DeviceAttr `json:"deviceattr"`
}

type DeviceAttrDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
