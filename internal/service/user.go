package service

import (
	"context"
	"strconv"

	"lottery-api/internal/model"
)

func (s *Svc) Join(c context.Context, req *model.UserInfoReq) (isOk bool, err error) {
	_, err = s.dao.SaveUsers(c, req.Phone, req.Article)
	if err != nil {
		s.dao.Logger.Printf("s.Join err(%v)", err)
		return
	}
	isOk = true
	return
}

func (s *Svc) VerifyUserPhone(c context.Context, phone int64) (isOk bool, messge string) {
	where := " where phone = " + strconv.FormatInt(phone, 10)
	user, err := s.dao.FindOneUser(c, where)
	if err != nil {
		s.dao.Logger.Printf("find user err (%v)", err)
	}
	if user != nil && user.Id > 0 {
		return false, "手机号已注册"
	}
	return true, "手机号未注册"
}

func (s *Svc) UserArticles(c context.Context) (data *model.UserInvolvesInfosReply, err error) {
	data = new(model.UserInvolvesInfosReply)
	data.List, err = s.dao.FetchUsers(c)
	if err != nil {
		s.dao.Logger.Printf("s.userArticles err(%v)", err)
	}
	return
}
