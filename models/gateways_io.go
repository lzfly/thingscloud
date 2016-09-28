package models

type GatewayPostForm struct {
	Gateway_sn      string `json:"gateway_sn"`
	Gw_model        string `json:"gw_model"`
	Gw_type         string `json:"gw_type"`
	Gw_mac          string `json:"gw_mac"`
	Wifi_ssid       string `json:"wifi_ssid"`
	Wifi_pwd        string `json:"wifi_pwd"`
	Hw_ver          string `json:"hw_ver"`
	Sw_ver          string `json:"sw_ver"`
	State           int32  `json:"state"`
	Operate_cap     int32  `json:"operate_cap"`
	Activetime      string `json:"activetime"`

}

type GatewayPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	GatewayInfo *Gateway `json:"gateway"`
}

type GatewayGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	GatewayInfo *Gateway `json:"gateway"`
}

type GatewayGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	GatewaysInfo []Gateway `json:"gateways"`
}

type GatewayPutForm struct {
	Gw_model        string `json:"gw_model"`
	Gw_type         string `json:"gw_type"`
	Gw_mac          string `json:"gw_mac"`
	Wifi_ssid       string `json:"wifi_ssid"`
	Wifi_pwd        string `json:"wifi_pwd"`
	Hw_ver          string `json:"hw_ver"`
	Sw_ver          string `json:"sw_ver"`
	State           int32  `json:"state"`
	Operate_cap     int32  `json:"operate_cap"`
}

type GatewayPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	GatewayInfo *Gateway `json:"gateway"`
}

type GatewayDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
