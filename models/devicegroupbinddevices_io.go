package models

type DeviceGroupBindDevicePostForm struct {
    Group_name           string `json:"group_name"`
	Device_sn            string `json:"device_sn"`
	Device_name          string `json:"device_name"`
    Type_code            int32  `json:"type_code"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           string `json:"attr_value"`
	Gateway_sn           string `json:"gateway_sn"`
	Username             string `json:"username"`
}

type DeviceGroupBindDevicePostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupBindDeviceInfo *DeviceGroupBindDevice `json:"devicegroupbinddevice"`
}

type DeviceGroupBindDeviceGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupBindDeviceInfo *DeviceGroupBindDevice `json:"devicegroupbinddevice"`
}

type DeviceGroupBindDeviceGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupBindDevicesInfo []DeviceGroupBindDevice `json:"devicegroupbinddevices"`
}

type DeviceGroupBindDevicePutForm struct {
    Group_name           string `json:"group_name"`
	Device_sn            string `json:"device_sn"`
	Device_name          string `json:"device_name"`
    Type_code            int32  `json:"type_code"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           string `json:"attr_value"`
	Gateway_sn           string `json:"gateway_sn"`
	Username             string `json:"username"`
}

type DeviceGroupBindDevicePutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupBindDeviceInfo *DeviceGroupBindDevice `json:"devicegroupbinddevice"`
}

type DeviceGroupBindDeviceDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
