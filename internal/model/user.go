package model

import "time"

type UserInfoReq struct {
	Phone      int64  `form:"phone" binding:"required,len=11"`
	VerifyID   string `form:"verify_id"`
	VerifyCode string `form:"verify_code" binding:"omitempty"`
	Article    string `form:"article" binding:"required,min=10,max=500"`
}

type UserInvolvesInfo struct {
	Id         int       `json:"id"`
	Phone      int64     `json:"phone"`
	DrawRight  int8      `json:"draw_right"`
	Article    string    `json:"article"`
	CreateTime time.Time `json:"create_time"`
}

type UserInvolvesInfosReply struct {
	List []*UserInvolvesInfo
}

const HaveDrawRight = 1
