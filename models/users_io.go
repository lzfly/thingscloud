package models

type UserPostForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	State     int32  `json:"state"`
}

type UserPostInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserInfo *User `json:"user"`
}

type UserGetOneInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserInfo *User `json:"user"`
}

type UserGetAllInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UsersInfo []User `json:"users"`
}

type UserPutForm struct {
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	State     int32  `json:"state"`
}

type UserPutInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
	UserInfo *User `json:"user"`
}

type UserDeleteInfo struct {
    Status   int    `json:"status"`
	Code     int    `json:"code"`
}
