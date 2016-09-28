package models

type DeviceGroupPostForm struct {
	Group_name      string `json:"group_name"`
	Username        string `json:"username"`
}

type DeviceGroupPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupInfo *DeviceGroup `json:"devicegroup"`
}

type DeviceGroupGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupInfo *DeviceGroup `json:"devicegroup"`
}

type DeviceGroupGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupsInfo []DeviceGroup `json:"devicegroups"`
}

type DeviceGroupPutForm struct {
	Group_name      string `json:"group_name"`
	Username        string `json:"username"`
}

type DeviceGroupPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceGroupInfo *DeviceGroup `json:"devicegroup"`
}

type DeviceGroupDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
