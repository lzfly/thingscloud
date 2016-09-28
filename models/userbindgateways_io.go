package models

type UserBindGatewayPostForm struct {
	Username           string  `json:"username"`
	Gateway_sn         string `json:"gateway_sn"`
    Is_master          int32  `json:"is_master"`
}

type UserBindGatewayPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindGatewayInfo *UserBindGateway `json:"userbindgateway"`
}

type UserBindGatewayGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindGatewayInfo *UserBindGateway `json:"userbindgateway"`
}

type UserBindGatewayGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindGatewaysInfo []UserBindGateway `json:"userbindgateways"`
}

type UserBindGatewayPutForm struct {
    Username           string  `json:"username"`
	Gateway_sn         string `json:"gateway_sn"`
    Is_master          int32  `json:"is_master"`
}

type UserBindGatewayPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserBindGatewayInfo *UserBindGateway `json:"userbindgateway"`
}

type UserBindGatewayDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
