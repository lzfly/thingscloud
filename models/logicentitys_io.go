package models

type LogicEntityPostForm struct {
	Username      string `json:"username"`
    Gateway_sn    string `json:"gateway_sn"`
	If_device_sn     string `json:"if_device_sn"`
	If_device_name     string `json:"if_device_name"`
	If_type_code     int32  `json:"if_type_code"`
	If_attr_code     int32  `json:"if_attr_code"`
	If_operate_code     string  `json:"if_operate_code"`
	If_attr_value     string `json:"if_attr_value"`
	If_attr_value2     string `json:"if_attr_value2"`
	Th_device_sn     string `json:"th_device_sn"`
	Th_device_name     string `json:"th_device_name"`
	Th_type_code     int32  `json:"th_type_code"`
	Th_attr_code     int32  `json:"th_attr_code"`
	Th_attr_value     string `json:"th_attr_value"`
}

type LogicEntityPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	LogicEntityInfo *LogicEntity `json:"logicentity"`
}

type LogicEntityGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	LogicEntityInfo *LogicEntity `json:"logicentity"`
}

type LogicEntityGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	LogicEntitysInfo []LogicEntity `json:"logicentitys"`
}

type LogicEntityPutForm struct {
	Username      string `json:"username"`
    Gateway_sn    string `json:"gateway_sn"`
	If_device_sn     string `json:"if_device_sn"`
	If_device_name     string `json:"if_device_name"`
	If_type_code     int32  `json:"if_type_code"`
	If_attr_code     int32  `json:"if_attr_code"`
	If_operate_code     string  `json:"if_operate_code"`
	If_attr_value     string `json:"if_attr_value"`
	If_attr_value2     string `json:"if_attr_value2"`
	Th_device_sn     string `json:"th_device_sn"`
	Th_device_name     string `json:"th_device_name"`
	Th_type_code     int32  `json:"th_type_code"`
	Th_attr_code     int32  `json:"th_attr_code"`
	Th_attr_value     string `json:"th_attr_value"`
}

type LogicEntityPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	LogicEntityInfo *LogicEntity `json:"logicentity"`
}

type LogicEntityDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
