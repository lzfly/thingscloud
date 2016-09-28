package models

type SessionPostForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type SessionPostInfo struct {
    Status   int     `json:"status"`
	Code     int     `json:"code"`
	Sessionid string `json:"sessionid"`
	Userid    int64  `json:"id"`
}

type SessionDeleteForm struct {
	Username  string `json:"username"`
}

type SessionDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}


