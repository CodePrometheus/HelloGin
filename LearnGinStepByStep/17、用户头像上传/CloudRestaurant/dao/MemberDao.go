package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/goes/logger"
	"github.com/matrix/go-matrix/log"
)

type MemberDao struct {
	*tool.Orm
}

//更新member记录，头像属性
func (md *MemberDao) UpdateMemberAvatar(userId int64, fileName string) int64 {
	member := model.Member{Avatar: fileName}
	result, err := md.Where(" id = ? ", userId).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//根据用户名和密码查询
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member

	password = tool.EncoderSha256(password)

	_, err := md.Where(" user_name = ? and password = ? ", name, password).Get(&member)
	if err != nil {
		log.ERROR(err.Error())
		return nil
	}

	return &member
}

//验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	md.Id()
	if _, err := md.Where(" phone = ? and code = ? ", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	if _, err := md.Where(" mobile  = ? ", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}
	return &member
}

//新用户的数据库插入操作
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
