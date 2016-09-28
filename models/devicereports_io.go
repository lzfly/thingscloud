package models

type DeviceReportPostForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           int64  `json:"attr_value"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceReportPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceReportInfo *DeviceReport `json:"devicereport"`
}

type DeviceReportGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceReportInfo *DeviceReport `json:"devicereport"`
}

type DeviceReportGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceReportsInfo []DeviceReport `json:"devicereports"`
}

type DeviceReportPutForm struct {
	Device_sn            string `json:"device_sn"`
	Attr_code            int32  `json:"attr_code"`
	Attr_value           int64  `json:"attr_value"`
	Gateway_sn           string `json:"gateway_sn"`
}

type DeviceReportPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	DeviceReportInfo *DeviceReport `json:"devicereport"`
}

type DeviceReportDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
