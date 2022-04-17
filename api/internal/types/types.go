// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `path:"userId"`
}

type UserInfoRes struct {
	Result   int8   `json:"result"`
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Sex      int32  `json:"sex"`
	City     int32  `json:"city"`
	SchoolId int32  `json:"schoolid"`
	Status   int8   `json:"status"`
	RegTime  string `json:"regTime"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}