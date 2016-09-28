package models

type DeviceTypePostForm struct {
	Type_name        string `json:"type_name"`
	Type_code        int32  `json:"type_code"`
}

type DeviceTypePostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceTypeInfo *DeviceType `json:"devicetype"`
}

type DeviceTypeGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceTypeInfo *DeviceType `json:"devicetype"`
}

type DeviceTypeGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceTypesInfo []DeviceType `json:"devicetypes"`
}

type DeviceTypePutForm struct {
	Type_name        string `json:"type_name"`
	Type_code        int32  `json:"type_code"`
}

type DeviceTypePutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceTypeInfo *DeviceType `json:"devicetype"`
}

type DeviceTypeDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}